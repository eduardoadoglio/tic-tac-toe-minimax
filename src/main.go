package main


func main() {
	gui := NewGUI()
	// baseInterface := setupBaseGameInterface(gameManager)
	baseInterface := gui.setupMenuInterface()
	gui.setWindowContent(baseInterface)
	gui.showContentAndRun()
}
