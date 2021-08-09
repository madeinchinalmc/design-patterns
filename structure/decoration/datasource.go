package decoration

import "fmt"

// 装饰可以改变组件接口所定义的操作。

type DataSource interface {
	WriteData(data string)
	ReadData() string
}

// 具体组件提供操作的默认实现。这些类在程序中可能会有几个变体。

type FileDataSource struct {
	FileName string
	Content  string
}

func NewFileDataSource(fileName string) *FileDataSource {
	return &FileDataSource{
		fileName,
		"",
	}
}

func (f *FileDataSource) WriteData(data string) {
	fmt.Printf("file %v data source write data %v \n ", f.FileName, data)
	f.Content = data
}

func (f *FileDataSource) ReadData() string {
	fmt.Printf("file %v data source read data \n", f.FileName)
	return f.Content
}

// 装饰基类和其他组件遵循相同的接口。该类的主要任务是定义所有具体装饰的封装接口。
// 封装的默认实现代码中可能会包含一个保存被封装组件的成员变量，并且对其进行初始化

type DataSourceDecorator struct {
	wrappee DataSource
}

func NewDataSourceDecorator(source DataSource) *DataSourceDecorator {
	var contentWrappee DataSource
	contentWrappee = &FileDataSource{}
	return &DataSourceDecorator{
		wrappee: contentWrappee,
	}
}

// 装饰基类会直接将所有工作分派给被封装组件。具体装饰中则可以新增一些额外的行为

func (d *DataSourceDecorator) WriteData(data string) {

	fmt.Printf("data source decorator write data %v  \n", data)
	d.wrappee.WriteData(data)
}

// 具体装饰可调用其父类的操作实现，而不是直接调用被封装的对象。

func (d *DataSourceDecorator) ReadData() string {
	fmt.Printf("data source decorator read data \n")
	return d.wrappee.ReadData()
}

// 具体装饰必须在被封装对象上调用方法，不过也可以自行在结果中添加一些内容。
// 装饰必须在调用封装对象之前或之后执行额外的行为。

type EncryptionDecorator struct {
	baseDecorator *DataSourceDecorator
}

func NewEncryptionDecorator(decorator *DataSourceDecorator) *EncryptionDecorator {
	return &EncryptionDecorator{
		baseDecorator: decorator,
	}
}

// 装饰基类会直接将所有工作分派给被封装组件。具体装饰中则可以新增一些额外的行为

func (e *EncryptionDecorator) WriteData(data string) {

	fmt.Printf("Encryption Decorator Encryption ing ... \n")
	e.baseDecorator.WriteData(data)
}

// 具体装饰可调用其父类的操作实现，而不是直接调用被封装的对象。

func (e *EncryptionDecorator) ReadData() string {
	fmt.Printf("Encryption Decorator Decryption ing ... \n")
	return e.baseDecorator.ReadData()
}

type CompressionDecorator struct {
	baseDecorator *DataSourceDecorator
}

func NewCompressionDecorator(decorator *DataSourceDecorator) *CompressionDecorator {
	return &CompressionDecorator{
		baseDecorator: decorator,
	}
}

// 装饰基类会直接将所有工作分派给被封装组件。具体装饰中则可以新增一些额外的行为

func (c *CompressionDecorator) WriteData(data string) {

	fmt.Printf("Compression Decorator Compression ing ... \n")
	c.baseDecorator.WriteData(data)
}

// 具体装饰可调用其父类的操作实现，而不是直接调用被封装的对象。

func (c *CompressionDecorator) ReadData() string {
	fmt.Printf("Compression Decorator UnCompression ing ... \n")
	return c.baseDecorator.ReadData()
}

// 客户端

type Application struct {
}

func (a *Application) DumbUsageExample() {
	var source DataSource
	source = NewFileDataSource("somefile.dat")
	source.WriteData("gogogo")

	datasourceDecorator := NewDataSourceDecorator(source)

	// 压缩
	compressionDecorator := NewCompressionDecorator(datasourceDecorator)
	compressionDecorator.WriteData("gogogo")

	// 加密
	encryptionDecorator := NewEncryptionDecorator(datasourceDecorator)
	encryptionDecorator.WriteData("gogogo")
}
