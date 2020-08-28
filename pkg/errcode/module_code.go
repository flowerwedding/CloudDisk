/**
 * @Title  module_code
 * @description  业务错误码
 * @Author  沈来
 * @Update  2020/8/5 16:09
 **/
package errcode

var (
	ErrorGetFileListFail  = NewError(10010, "获取文件列表失败")
	ErrorGetFileFail      = NewError(10011, "获取文件列表失败")
	ErrorCreateFileFail   = NewError(10012, "创建文件失败")
	ErrorCreateFileQtFail = NewError(10013, "创建文件二维码失败")
	ErrorDownFileFail     = NewError(10014, "创建文件失败")
	ErrorDeleteFileFail   = NewError(10015, "删除文件失败")
	ErrorCountFileFail    = NewError(10016, "统计文件失败")
)
