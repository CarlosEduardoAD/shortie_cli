package cloudconverter

type ICloudConverterRepository interface {
	ConvertFile(file string, extension string, path string) error
}

type CloudConverterRepository struct {
	url string
}

func NewCloudConverterRepository(url string) *CloudConverterRepository {
	return &CloudConverterRepository{url: url}
}

func (c *CloudConverterRepository) ConvertFile(file string, extension string, path string) error {
	return nil
}
