package model

type BaseLogInfo struct {
	FilePath     string `json:"file_path,omitempty"`
	CurrentLine  int    `json:"current_line,omitempty"`
	FuncName     string `json:"func_name,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
