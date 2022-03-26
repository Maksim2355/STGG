package tmplengine

import "stgg/crossplatform"

const TemplateDirName = "templates"

const TemplatesPath = "." + crossplatform.PATH_SEPARATOR + TemplateDirName

const GlobalVariablesFileName = "globalVariables.yaml"

const GlobalVariablesPath = "." + crossplatform.PATH_SEPARATOR + GlobalVariablesFileName

const GeneratedFilesDirName = "temp"

const PathGenerated = "." + crossplatform.PATH_SEPARATOR + GeneratedFilesDirName
