package merchants

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"

	"github.com/cobuildlab/pex-cmd/utils"
	"github.com/spf13/cobra"
)

//CmdUploadList Command to list the available merchants files to upload
var CmdUploadList = &cobra.Command{
	Use:   "list",
	Short: "List of merchants files available to upload to the database",
	Long:  "List of merchants files available to upload to the database",
	Run: func(cmd *cobra.Command, args []string) {
		fileList, err := UploadList()
		if err != nil {
			log.Println(err)
			return
		}

		var count uint64
		var totalSize int64
		for _, v := range fileList {
			count++
			totalSize += v.Size
			log.Println("╭─Filename:", v.Name)
			log.Println("├─⇢ Size:", v.Size)
			log.Println("╰─⇢ Time:", v.ModTime)
			log.Println()
		}
		log.Println("╭─Total files:", count)
		log.Println("╰─Total size:", totalSize)
	},
}

//UploadList List the available merchants files to upload
func UploadList() (filesList []utils.FileInfo, err error) {
	files, err := ioutil.ReadDir(DecompressPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if ext == MerchantFileFormatExt {
				var fileXML utils.FileInfo
				fileXML.Name = file.Name()
				fileXML.Size = file.Size()
				fileXML.ModTime = file.ModTime()

				filesList = append(filesList, fileXML)
			}
		}
	}

	sort.Slice(filesList, func(i, j int) bool {
		return filesList[i].Size > filesList[j].Size
	})

	return
}
