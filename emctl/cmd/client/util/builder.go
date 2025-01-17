/*
 * Copyright (c) 2017, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// FileExtensions describe that what's the extended file name of the EaseMesh configuration file should have
var FileExtensions = []string{".json", ".yaml", ".yml"}

type (
	// VisitorBuilder is a builder that build a visitor to visit func
	VisitorBuilder interface {
		HTTPAttemptCount(httpGetAttempts int) VisitorBuilder
		FilenameParam(filenameOptions *FilenameOptions) VisitorBuilder
		CommandParam(commandOptions *CommandOptions) VisitorBuilder
		Command() VisitorBuilder
		Do() ([]Visitor, error)
		File() VisitorBuilder
		URL(httpAttemptCount int, urls ...*url.URL) VisitorBuilder
		Stdin() VisitorBuilder
	}
	visitorBuilder struct {
		visitors          []Visitor
		decoder           Decoder
		httpGetAttempts   int
		errs              []error
		singleItemImplied bool
		commandOptions    *CommandOptions
		filenameOptions   *FilenameOptions
		stdinInUse        bool
	}

	// CommandOptions holds command option
	CommandOptions struct {
		// Kind is required.
		Kind string
		// Name is allowed to be empty.
		Name string
	}

	// FilenameOptions holds filename option
	FilenameOptions struct {
		Filenames []string
		Recursive bool
	}
)

// NewVisitorBuilder return a VisitorBuilder
func NewVisitorBuilder() VisitorBuilder {
	return &visitorBuilder{httpGetAttempts: 3, decoder: newDefaultDecoder()}
}

func (b *visitorBuilder) HTTPAttemptCount(httpGetAttempts int) VisitorBuilder {
	b.httpGetAttempts = httpGetAttempts
	return b
}

func (b *visitorBuilder) FilenameParam(filenameOptions *FilenameOptions) VisitorBuilder {
	b.filenameOptions = filenameOptions
	return b
}

func (b *visitorBuilder) CommandParam(commandOptions *CommandOptions) VisitorBuilder {
	b.commandOptions = commandOptions
	return b
}

func (b *visitorBuilder) Command() VisitorBuilder {
	if b.commandOptions == nil {
		return b
	}

	b.visitors = append(b.visitors, newCommandVisitor(
		b.commandOptions.Kind,
		b.commandOptions.Name,
	))

	return b
}

func (b *visitorBuilder) Do() ([]Visitor, error) {
	b.Command()
	b.File()

	if len(b.errs) != 0 {
		return nil, fmt.Errorf("%+v", b.errs)
	}

	return b.visitors, nil
}

func (b *visitorBuilder) File() VisitorBuilder {
	if b.filenameOptions == nil {
		return b
	}

	recursive := b.filenameOptions.Recursive
	paths := b.filenameOptions.Filenames
	for _, s := range paths {
		switch {
		case s == "-":
			b.Stdin()
		case strings.Index(s, "http://") == 0 || strings.Index(s, "https://") == 0:
			url, err := url.Parse(s)
			if err != nil {
				b.errs = append(b.errs, fmt.Errorf("the URL passed to filename %q is not valid: %v", s, err))
				continue
			}
			b.URL(b.httpGetAttempts, url)
		default:
			if !recursive {
				b.singleItemImplied = true
			}
			b.path(recursive, s)
		}
	}

	return b
}

func (b *visitorBuilder) URL(httpAttemptCount int, urls ...*url.URL) VisitorBuilder {
	for _, u := range urls {
		b.visitors = append(b.visitors, &urlVisitor{
			URL:              u,
			streamVisitor:    newStreamVisitor(nil, b.decoder, u.String()),
			HTTPAttemptCount: httpAttemptCount,
		})
	}
	return b
}

func (b *visitorBuilder) Stdin() VisitorBuilder {
	if b.stdinInUse {
		b.errs = append(b.errs, errors.Errorf("Stdin already in used"))
	}
	b.stdinInUse = true
	b.visitors = append(b.visitors, FileVisitorForSTDIN(b.decoder))
	return b
}

func (b *visitorBuilder) path(recursive bool, paths ...string) VisitorBuilder {
	for _, p := range paths {
		_, err := os.Stat(p)
		if os.IsNotExist(err) {
			b.errs = append(b.errs, fmt.Errorf("the path %q does not exist", p))
			continue
		}
		if err != nil {
			b.errs = append(b.errs, fmt.Errorf("the path %q cannot be accessed: %v", p, err))
			continue
		}

		visitors, err := expandPathsToFileVisitors(b.decoder, p, recursive, FileExtensions)
		if err != nil {
			b.errs = append(b.errs, fmt.Errorf("error reading %q: %v", p, err))
		}

		b.visitors = append(b.visitors, visitors...)
	}
	if len(b.visitors) == 0 {
		b.errs = append(b.errs, fmt.Errorf("error reading %v: recognized file extensions are %v", paths, FileExtensions))
	}
	return b
}
