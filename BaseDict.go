package util

//为map[int]int对像绑定的一些方法
type BaseDcit struct {
	Data map[int]int
}

//Clone 复制一个BaseData
func (this BaseDcit) Clone() *BaseDcit {
	result := new(BaseDcit)
	result.Data = make(map[int]int)
	for k, v := range this.Data {
		result.Data[k] = v
	}
	return result
}

//NewBaseDictString 用字符串初始化一个数据
func NewBaseDictString(str *String) *BaseDcit {
	result := new(BaseDcit)
	result.Data = make(map[int]int)
	if str == nil {
		return result
	}
	str = str.Replace("|", ";")
	d := str.ToString()
	if d == "" || d == "0" {
		return result
	}
	arr := StringToIntArray(d, ";")
	for i := 0; i < len(arr); i += 2 {
		result.Data[arr[i]] = arr[i+1]
	}
	return result
}

//UpData 更新指定数据
func (this *BaseDcit) UpData(key, num int) {
	v, _ := this.Data[key]
	if num+v > 0 {
		this.Data[key] = v + num
	} else {
		this.Data[key] = v + num
	}
}

//UpDataBc批量用别的数据，更新本数据
func (this *BaseDcit) UpDataBc(addbc, delbc *BaseDcit) {
	if delbc != nil {
		for k, n := range delbc.Data {
			this.UpData(k, -n)
		}
	}
	if addbc != nil {
		for k, n := range addbc.Data {
			this.UpData(k, n)
		}
	}
}

//GetNumByKey指定数据的值
func (this *BaseDcit) GetNumByKey(key int) int {
	v, ok := this.Data[key]
	if !ok {
		return 0
	}
	return v
}

//ToString 字符串化
func (this *BaseDcit) ToString() string {
	sb := NewStringBuilder()
	t := 0
	for k, v := range this.Data {
		if t == 0 {
			t++
		} else {
			sb.Append(";")
		}
		sb.AppendInt(k)
		sb.Append(";")
		sb.AppendInt(v)
	}
	return sb.ToString()
}

//Count 总数量
func (this *BaseDcit) Count() (result int) {
	for _, n := range this.Data {
		result += n
	}
	return result
}

//MaxItem 最大数值的KEY，value
func (this *BaseDcit) MaxItem() (key, num int) {
	for k, n := range this.Data {
		if n > num {
			key, num = k, n
		}
	}
	return key, num
}

//Clear清数据
func (this *BaseDcit) Clear() {
	this.Data = make(map[int]int)
}
