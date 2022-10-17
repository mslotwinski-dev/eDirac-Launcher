package generate

import (
	os "os"
	strconv "strconv"
)

func Generate(basicpath string, i int) {
	var path = basicpath + "/" + strconv.Itoa(i)
	os.MkdirAll(path, os.ModePerm)

	// CORE FILE
	os.WriteFile(basicpath+"/"+strconv.Itoa(i)+".ts", []byte(PlaceholderCore), 0644)

	os.MkdirAll(path+"/img", os.ModePerm)

	path += "/content/"
	os.MkdirAll(path, os.ModePerm)
	os.WriteFile(path+"/Idea.vue", []byte(PlaceholderIdea), 0644)
	os.WriteFile(path+"/Index.vue", []byte(PlaceholderVIP), 0644)
	os.WriteFile(path+"/Author.vue", []byte(PlaceholderVIP), 0644)
	os.WriteFile(path+"/Bibliography.vue", []byte(PlaceholderIdea), 0644)

	os.MkdirAll(path+"Wprowadzenie/", os.ModePerm)
	os.MkdirAll(path+"Wprowadzenie/1/", os.ModePerm)
	os.WriteFile(path+"Wprowadzenie/1/1.vue", []byte(PlaceholderFile), 0644)

}
