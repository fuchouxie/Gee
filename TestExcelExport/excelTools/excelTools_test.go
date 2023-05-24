package excelTools

import (
	"fmt"
	"testing"
)

func TestExportToPath(t *testing.T) {
	header := make([]map[string]string, 0)
	header = append(header, map[string]string{"key": "id", "title": "ID", "width": "10", "is_num": "0"})
	header = append(header, map[string]string{"key": "settlement_num", "title": "结算编号", "width": "20", "is_num": "0"})
	header = append(header, map[string]string{"key": "settlement_status", "title": "状态", "width": "15", "is_num": "0"})
	header = append(header, map[string]string{"key": "enterprise_name", "title": "企业客户", "width": "20", "is_num": "0"})
	header = append(header, map[string]string{"key": "settlement_amount", "title": "结算金额", "width": "15", "is_num": "0"})
	header = append(header, map[string]string{"key": "deduction_amount", "title": "预付款抵扣金额", "width": "15", "is_num": "0"})
	header = append(header, map[string]string{"key": "unpaid_amount", "title": "应回款金额", "width": "10", "is_num": "0"})
	header = append(header, map[string]string{"key": "apply_time", "title": "申请时间", "width": "15", "is_num": "0"})
	header = append(header, map[string]string{"key": "order_num", "title": "关联订单数", "width": "20", "is_num": "0"})
	header = append(header, map[string]string{"key": "operator_name", "title": "操作人员", "width": "20", "is_num": "0"})
	data := make([]map[string]interface{}, 0)
	data = append(data, map[string]interface{}{
		"id":                1,
		"settlement_num":    "100122102117188572341",
		"settlement_status": "结算关闭",
		"enterprise_name":   "公司A",
		"settlement_amount": "100000",
		"deduction_amount":  "40000",
		"unpaid_amount":     "60000",
		"apply_time":        "2022-11-24  7:00",
		"order_num":         25,
		"operator_name":     "liudehua",
	})

	header1 := make([]map[string]string, 0)
	header1 = append(header1, map[string]string{"key": "id", "title": "ID", "width": "10", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "order_serial_no", "title": "线下订单号", "width": "20", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "hi_order_no", "title": "订单号", "width": "15", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "hi_appointment_id", "title": "预约单号", "width": "20", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "user_name", "title": "用户名", "width": "15", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "user_phone", "title": "手机号", "width": "15", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "user_gender_name", "title": "性别", "width": "10", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "user_birth", "title": "生日", "width": "15", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "credential_type_name", "title": "证件类型", "width": "20", "is_num": "0"})
	header1 = append(header1, map[string]string{"key": "credential_no", "title": "证件号", "width": "20", "is_num": "0"})
	data1 := make([]map[string]interface{}, 0)
	data1 = append(data1, map[string]interface{}{
		"id":                   1,
		"order_serial_no":      "100122102117188572341",
		"hi_order_no":          "100122102117188572341",
		"hi_appointment_id":    "979879879879",
		"user_name":            "xfc",
		"user_phone":           "15015010501",
		"user_gender_name":     "男",
		"user_birth":           "2022-11-24",
		"credential_type_name": "身份证",
		"credential_no":        "100122102117188572341",
	})

	sheetModels := []sheetModel{}
	sheetModels = append(sheetModels, sheetModel{
		sheetName: "结算",
		title:     header,
		data:      data,
	})
	sheetModels = append(sheetModels, sheetModel{
		sheetName: "订单",
		title:     header1,
		data:      data1,
	})
	sheetModels = append(sheetModels, sheetModel{
		sheetName: "订单",
		title:     header1,
		data:      data1,
	})
	sheetModels = append(sheetModels, sheetModel{
		sheetName: "订单",
		title:     header1,
		data:      data1,
	})
	excel := NewMyExcel()
	path, err := excel.ExportToPath(sheetModels, "D://", "测试导出")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(path)
}
