package main


func main() {
	gui := NewGUI()
	baseInterface := gui.setupMenuInterface()
	gui.setWindowContent(baseInterface)
	gui.showContentAndRun()
}
