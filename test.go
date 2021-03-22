func setup() {
	//init 配置
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}


func TestDemo(t *testing.T) {
  fmt.PrintLn("demo")
}
