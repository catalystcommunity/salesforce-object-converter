package internal

const ProtoAny = "google.protobuf.Any"
const ProtoBase64 = ProtoAny
const ProtoBoolean = "bool"
const ProtoCombobox = ProtoString
const ProtoCurrency = ProtoDouble
const ProtoDataCategoryGroupReference = ProtoString
const ProtoDate = ProtoString
const ProtoDatetime = ProtoString
const ProtoDouble = "double"
const ProtoEmail = ProtoString
const ProtoEncryptedString = ProtoString
const ProtoId = ProtoString
const ProtoInt = "int64"
const ProtoMultiPickList = "repeated string"
const ProtoPercent = ProtoDouble
const ProtoPhone = ProtoString
const ProtoPicklist = ProtoString
const ProtoReference = ProtoString
const ProtoString = "string"
const ProtoTextArea = ProtoString
const ProtoTime = ProtoString
const ProtoUrl = ProtoString

var SfProtoTypeMap = map[string]string{
	SfAny:                        ProtoAny,
	SfBase64:                     ProtoBase64,
	SfBoolean:                    ProtoBoolean,
	SfCombobox:                   ProtoCombobox,
	SfCurrency:                   ProtoCurrency,
	SfDataCategoryGroupReference: ProtoDataCategoryGroupReference,
	SfDate:                       ProtoDate,
	SfDatetime:                   ProtoDatetime,
	SfDouble:                     ProtoDouble,
	SfEmail:                      ProtoEmail,
	SfEncryptedString:            ProtoEncryptedString,
	SfId:                         ProtoId,
	SfInt:                        ProtoInt,
	SfMultiPickList:              ProtoMultiPickList,
	SfPercent:                    ProtoPercent,
	SfPhone:                      ProtoPhone,
	SfPicklist:                   ProtoPicklist,
	SfReference:                  ProtoReference,
	SfString:                     ProtoString,
	SfTextArea:                   ProtoTextArea,
	SfTime:                       ProtoTime,
	SfUrl:                        ProtoUrl,
}
