package sqtobfuscate

type testSet struct {
	key    string
	plain  string
	cipher string
}

var sqtTestSets = []testSet{
	{
		key:    "7URzs0Ee/TAIhPUqlexG/A==",
		plain:  "{\"method\":\"staff.batch.add\",\"ts\":1641264851,\"entId\":\"39597\",\"staffInfos\":[{\"name\":\"李杰\",\"entStaffNum\":\"D0920150228\"}]}",
		cipher: "38tECwETiBF6vWI30yJIASRJqsQGB4dxcjElrq6cT4L2VwfGlA17jv5gasYObDlc5UYF5Rl2MpRQ44x_wZ9uCnKMlhY21dQHSHrHPfrQ4bwl8SM7uPCiRSy-mXQmI4Ea_DUDHCeSRAX_a05CY-V8vSoQas9PkoBHNJZw7ylxwik",
	},
	{
		key:    "xd1nzb/N9Nx3+VoImzCsnw==",
		plain:  "{\"sign\":\"sgW1bxc7oatFhOJXAeHnNg==\",\"ts\":1512964057,\"method\":\"waimai.poi.list\",\"longitude\":116488645,\"latitude\":40007069}",
		cipher: "UgJn07uNgW7S7fJK0R0xVbaLxoCGPQIzoP-_K4Hmp4RduGszhm2mbUs2toZhCtXKP5JGXVTZ9kGts2Wx3IJQCd90ptMoJTDB0vu7mkedEr4KZCvZn77EZLssMC5SpXilmQ-5RXHzvMIT0ASH-IXepTP_O16U37QqCkEb5L1WLy4",
	},
}
