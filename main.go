package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Todo struct {
	task    string
	endDate string
}

var todo_list []Todo

var app = tview.NewApplication()

var text = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("(a) to add a new contact \n(q) to quit")

var form = tview.NewForm()
var pages = tview.NewPages()

var TodoList = tview.NewList().ShowSecondaryText(false)

var flex = tview.NewFlex()

func addTodoForm() {
	var todo Todo = Todo{}

	form.AddInputField("Task", "", 100, nil, func(text string) {
		todo.task = text
	})

	form.AddInputField("Date", "", 10, nil, func(text string) {
		todo.endDate = text
	})
	form.AddButton("Save", func() {
		todo_list = append(todo_list, todo)
		addTodoList()
		pages.SwitchToPage("Menu")
	})
}

func addTodoList() {
	for index, todo := range todo_list {
		TodoList.AddItem(todo.task, "", rune(49+index), nil)
	}
}

func main() {
	flex.SetDirection(tview.FlexRow).
		AddItem(TodoList, 0, 1, true).
		AddItem(text, 0, 1, false)

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add Form", form, true, false)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			form.Clear(true)
			addTodoForm()
			pages.SwitchToPage("Add Form")
		}
		return event
	})
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
