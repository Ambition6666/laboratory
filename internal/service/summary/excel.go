package summary

import (
	"fmt"
	"laboratory/internal/dao"
	"laboratory/log"
	"laboratory/pkg/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

// 获取excel表预约信息
func GetBookingINFOExcel(lid string) (int, string, any) {
	data, err := dao.SearchAppointment(lid)
	if err != nil {
		log.SugarLogger.Error("获取预约信息失败", err)
		return http.StatusInternalServerError, "导出失败", nil
	}

	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			log.SugarLogger.Error("关闭excel出错", err)
		}
	}()

	// 创建一个工作表
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		log.SugarLogger.Error("创建excel工作表失败", err)
		return http.StatusInternalServerError, "导出失败", nil
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// 按行赋值
	for i := 1; i <= len(data); i++ {
		s := fmt.Sprintf("A%d", i)
		log.SugarLogger.Debugln(s)
		arr := utils.StoArr(data[i-1])
		log.SugarLogger.Debugln(arr)
		err = f.SetSheetRow("Sheet1", s, &arr)
		if err != nil {
			log.SugarLogger.Error("创建excel工作表失败", err)
			return http.StatusInternalServerError, "导出失败", nil
		}
	}

	// 工作路径
	workdir, _ := os.Getwd()

	name :=  fmt.Sprintf("%s.xlsx", lid)
	path := filepath.Join(workdir, "excel", name)

	// 保存excel
	if err := f.SaveAs(path); err != nil {
		log.SugarLogger.Error("工作表保存失败", err)
		return http.StatusInternalServerError, "导出失败", nil
	}

	return http.StatusOK, "导出成功", name
}
