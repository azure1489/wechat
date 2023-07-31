package model

type GetFileCrc32Req struct {
	Filepath string `json:"filepath"`
}

type GetFileCrc32Result struct {
	GetFileCrc32 string `json:"Get_File_crc32"`
}
