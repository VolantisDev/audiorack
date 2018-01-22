# VST Guide

VST Guide is a tool for audio producers to help them manage their audio
asset library.

## Supported Formats

- VST instruments and effects (VST2 and VST3)
- Kontakt instruments
- Reaktor instruments and effects
- UVI instruments
- Samples (Multi-format)
- Plugin presets

## Future Development

In the future, VST Guide will also support AU, AAX, RTAS, and
potentially other types of standardized formats.

Feel free to submit bug reports and feature requests, or to implement
the functionality yourself and submit a pull request. We're grateful for
any contributions!

## Repository

This repository consists of a mono-repo of a number of sub-projects
which are (or can be) developed independently. The main reason the
primary components are all combined into one repository is to allow
people to understand and execute the overall application in a simple
and consistent way. It also allows people to get up to speed developing
the application easily, and be more productive when working on multiple
components at once.

## Languages

Parts of VST Guide are written with Go, while the frontend is written
as an Electron application using the Quasar framework. The frontend
and backend are integrated using `go-astelectron` to allow for seamless
communication between the two.

Due to Go being the primary backend language, Node.js is only used in
the Vue.js components themselves (meaning it gets compiled to standard
browser JS). `electron-remote` can be called to access backend
functionality in the same way as you would expect from a pure Electron
app.

Babel is used to compile down from ES6, and thus ES6 should be used for
any and all Javascript written for the application to keep things
pleasant to write and maintain.
