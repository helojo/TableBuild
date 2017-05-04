package xml

import (
	"TableBuild/generator"
	"TableBuild/tool"
	"strings"
)

type Xml_generator struct {
	M_cfgbase *generator.Cfgbase_t

	M_xml_dir string
}

func (x *Xml_generator) Generate() {
	if false == tool.Exist(&x.M_xml_dir) {
		tool.EchoErr("错误: <%s>路径不存在", x.M_xml_dir)
		return
	}

	x.save_as_xml()
}

func (x *Xml_generator) save_as_xml() {
	// 将excel中各个工作表数据导出到对应的xml文件中
	for _, v := range x.M_cfgbase.Cfgs {
		var xml string = x.M_xml_dir + "\\" + v.En_name + ".xml"

		var xml_text string
		get_xml_text(&v, &xml_text)

		tool.WriteFile(&xml, &xml_text)
		tool.EchoOk("导出[%s] -> [%s]成功", v.Cn_name, xml)
	}
}

// 转码:
//     1. 引号 -> &quot;
func escape_xml(str *string) {
	*str = strings.Replace(*str, "\"", "&quot;", -1)
}

// 将数据拼接成xml格式
func get_xml_text(cfg *generator.Cfg_t, xml_text *string) {
	*xml_text = "<" + cfg.En_name + " xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\">\n"
	var row_text string

	var n_field int = len(cfg.Fields)
	for _, row := range cfg.Table {
		// 将数据拼接成形如<row a="a1" b="b1" c="c1"....>的xml文本
		row_text = "	<row"

		for c := 0; c < n_field; c++ {
			escape_xml(&row[c])
			row_text += " " + cfg.Fields[c].En_name + "=\"" + row[c] + "\""
		}

		*xml_text += row_text + "/>\n"
	}

	*xml_text += "</" + cfg.En_name + ">"
}
