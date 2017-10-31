// trainctl is an opinionated tool for managing training materials.
//
// trainctl allows you to mix multiple training modules in various combinations,
// creating courses from the modules.
//
// Modules are self-contained topics.  A module might represent a single presentation
// at a conference, or a single topic of a multi-topic course.
//
// Courses are a collection of modules assembled in a specific order to create a logical
// progression of learning material.
//
// Modules are stored in a base folder.  We'll refer to the base folder as $MODULEDIR
// Each module consists of a present slide file, a json manifest and a subdirectory
// bearing the same name as the module, containing
// the slide's assets, in `images` and `includes` directories.
//
// Source code intended for student consumption is stored in `$MODULEDIR/src/modulename`.
// The expectation is that each course will be a self-contained $GOPATH, so the import
// names of each module are simply the modulename.
//
// trainctl is a tool that allows you to manage modules individually, with the end goal of
// assembling them into a course which will be presented/taught.
//
// trainctl uses a configuration file stored at $HOME/.trainctl.yaml which can be created using
// `trainctl createconfig`.  It is not required, but will pre-fill many of the fields needed for
// every command, so it is strongly recommended.
//
// Usage:
//  trainctl [command]
//
// Available Commands:
//  assemble     Assemble modules into a course
//  create       Create a new module skeleton
//  createconfig Create a config file in your home directory
//  info         Get info about a module
//  package      package a course into a downloadable zip file
//  related      Add related material to a module
//  remove       Remove a new module skeleton
//  search       Search modules by metadata
//
// Start with `trainctl createconfig` to create a configuration file in your home directory.
// Edit this file with real values representing your configuration.
// The $MODULEDIR should be a source controlled directory.
// The $COURSEDIR can be any directory, source control not required (or recommended).
//
// To create a module run:
//  trainctl create --name mymodule
//
// This creates the following files and directories:
//  $MODULEDIR/mymodule.slide
//  $MODULEDIR/mymodule.json
//  $MODULEDIR/mymodule/images
//  $MODULEDIR/mymodule/includes
// These files represent the module's slide and assets required to present the slide.
//
// Additionally, the following directories are created:
//  $MODULEDIR/src/mymodule/demos
//  $MODULEDIR/src/mymodule/exercises
//  $MODULEDIR/src/mymodule/solutions
// These directories are intended to be a self-contained $GOPATH for the module, and eventually
// the courses that contain this module.
//
// Create one or more modules this way, then assemble them into a course with the
// `trainctl assemble` command like this:
//  trainctl assemble --course "My Class" --modules module1,module2,module4 --shortname myclass
// The assemble command will create a course in the $COURSEDIR folder with the name `myclass` in the
// above example.  The following files are created
//  $COURSEDIR/mymodule.json
//  $COURSEDIR/mymodule.slide
// The .json file is a manifest of the course's contents.
// The .slide file is a symlink to your module's .slide file, with the addition of a numbered
// prefix, establishing the order the modules you specified in the `assemble` command.
// The module directory is symlinked to:
//  $COURSEDIR/mymodule
// with the subdirectories intact so the slide's assets will still appear on the slide.
//
// The `package` sub-command takes a course from your $COURSEDIR and creates a zipfile from
// it, which has the effect of de-referencing the symlinks.  The benefit if this approach is that
// your courses will always contain the latest up-to-date changes from the modules included in them,
// up to the point where you memorialize them immediately before presenting your class.
// Repackaging the course in the future will update the zipfile with files in the symlinked
// modules that may have changed since last giving the course.
//
//  trainctl package --shortname mycourse
// This creates a zipfile called mycourse.zip with the contents of the `mycourse` included.
// This is the course material you will use to present, or give to your students.
// It also includes a Vagrantfile and a supporting bootstrap bash script which provides
// a fully operational Go development with the course's /src shared into the Vagrant container.
//
// The `search` and `info` sub-commands give you information about modules in your
// $MODULEDIR so that you can choose which modules to assemble into a course.
package main
