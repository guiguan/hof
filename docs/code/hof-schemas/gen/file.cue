package gen

// A file which should be generated by hof
#File: {
	// The local input data, any struct
	// The Generator.In will be added here
	//   but will not replace any values set locally
	In?: {...} // for templates
	Val?: {...} // for datafiles

	// The full path under the output location
	// empty implies don't generate, even though it may end up in the out list
	Filepath?: string

	// Only one of these three may be set
	// The template contents
	TemplateContent?: string
	// Path into the loaded templates
	TemplatePath?: string
	// Writes a datafile, bypassing template rendering
	DatafileFormat?: "cue" | "json" | "yaml" | "xml" | "toml"

	// TODO, we would like to make the above a disjunction (multi-field)
	// but it results in a significant slowdown 50-100% for hof self-gen
	// Most likely need to wait for structural sharing to land in cue

	// Alternative Template Delimiters
	Delims: #TemplateDelims
	TemplateDelims?: Delims

	// Formatting Control
	Formatting?: {
		Disabled?: bool
		// Name of the formatter, like 'prettier' or 'black'
		Formatter: string
		// formatter specific configuration
		Config: _
	}

	// Note, intentionally closed to prevent user error when creating GenFiles
}

// deprecated
#HofGeneratorFile: #File

