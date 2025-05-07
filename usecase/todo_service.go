package usecase

import (
	"sync"
	"time"
	"todoapp/domain/model"
)

var (
	todolists = []*model.TodoList{
		{
			ID:         1,
			Title:      "Alışveriş Listesi",
			Percentage: 50,
			CreatedAt:  time.Now().Add(-48 * time.Hour),
			ModifiedAt: time.Now().Add(-24 * time.Hour),
			DeletedAt:  time.Time{},
		},
		{
			ID:         2,
			Title:      "Yazılım Projesi",
			Percentage: 25,
			CreatedAt:  time.Now().Add(-72 * time.Hour),
			ModifiedAt: time.Now().Add(-48 * time.Hour),
			DeletedAt:  time.Time{},
		},
		{
			ID:         3,
			Title:      "Tatilde Yapılacaklar",
			Percentage: 0,
			CreatedAt:  time.Now().Add(-96 * time.Hour),
			ModifiedAt: time.Now().Add(-90 * time.Hour),
			DeletedAt:  time.Time{},
		},
		{
			ID:         4,
			Title:      "Okul Dönemi Hazırlıkları",
			Percentage: 75,
			CreatedAt:  time.Now().Add(-24 * time.Hour),
			ModifiedAt: time.Now(),
			DeletedAt:  time.Time{},
		},
		{
			ID:         5,
			Title:      "Günlük Rutin",
			Percentage: 100,
			CreatedAt:  time.Now().Add(-10 * 24 * time.Hour),
			ModifiedAt: time.Now().Add(-8 * 24 * time.Hour),
			DeletedAt:  time.Time{},
		},
	}
	todos = []*model.TodoStep{{ID: 1, TodoListId: 1, Message: "Süt al", IsDone: true, CreatedAt: time.Now().Add(-47 * time.Hour)},
		{ID: 2, TodoListId: 1, Message: "Yumurta al", IsDone: true, CreatedAt: time.Now().Add(-46 * time.Hour)},
		{ID: 3, TodoListId: 1, Message: "Un al", IsDone: false, CreatedAt: time.Now().Add(-45 * time.Hour)},

		{ID: 4, TodoListId: 2, Message: "Git reposu oluştur", IsDone: true, CreatedAt: time.Now().Add(-70 * time.Hour)},
		{ID: 5, TodoListId: 2, Message: "Proje yapısını oluştur", IsDone: false, CreatedAt: time.Now().Add(-69 * time.Hour)},
		{ID: 6, TodoListId: 2, Message: "API endpoint'lerini yaz", IsDone: false, CreatedAt: time.Now().Add(-68 * time.Hour)},

		{ID: 7, TodoListId: 3, Message: "Kitap oku", IsDone: false, CreatedAt: time.Now().Add(-94 * time.Hour)},
		{ID: 8, TodoListId: 3, Message: "Yüzmeye git", IsDone: false, CreatedAt: time.Now().Add(-93 * time.Hour)},

		{ID: 9, TodoListId: 4, Message: "Ders kitaplarını al", IsDone: true, CreatedAt: time.Now().Add(-20 * time.Hour)},
		{ID: 10, TodoListId: 4, Message: "Kalem & defter", IsDone: true, CreatedAt: time.Now().Add(-19 * time.Hour)},
		{ID: 11, TodoListId: 4, Message: "Kayıt kontrolü yap", IsDone: false, CreatedAt: time.Now().Add(-18 * time.Hour)},

		{ID: 12, TodoListId: 5, Message: "Diş fırçala", IsDone: true, CreatedAt: time.Now().Add(-9 * 24 * time.Hour)},
		{ID: 13, TodoListId: 5, Message: "Sabah yürüyüşü", IsDone: true, CreatedAt: time.Now().Add(-9 * 24 * time.Hour)},
		{ID: 14, TodoListId: 5, Message: "Günlük yaz", IsDone: true, CreatedAt: time.Now().Add(-9 * 24 * time.Hour)}}
	nextTodoId     = len(todos)
	nextTodoListId = len(todolists)
	mu             sync.Mutex
	muLock         bool
)

func MuUnlock() {
	mu.Unlock()
	muLock = false
}
func GetTodosByUserId(userId int) []*model.TodoStep {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}

	user := GetUserById(userId)

	var todoSteps []*model.TodoStep
	for _, list := range user.TodoLists {
		if steps := GetTodosByListId(list.ID); steps != nil {
			todoSteps = append(todoSteps, steps...)

		}

	}

	return todoSteps
}
func AllTodoLists() []*model.TodoList {
	return todolists
}
func AllTodos() []*model.TodoStep {
	return todos
}
func AddTodo(todo *model.TodoStep) {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}
	todo.ID = nextTodoId
	nextTodoId++
	todo.CreatedAt = time.Now()
	todos = append(todos, todo)
}
func CreateTodoList(todolist *model.TodoList, userId int) {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}
	todolist.ID = nextTodoListId
	todolist.CreatedAt = time.Now()
	nextTodoListId++
	user := GetUserById(userId)
	user.TodoLists = append(user.TodoLists, todolist)
	todolists = append(todolists, todolist)

}
func ChangeMessage(todoId int, newmsg string) {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}
	todo := GetTodoById(todoId)
	if todo == nil {
		return
	}
	todo.Message = newmsg
	todo.ModifiedAt = time.Now()
}
func GetTodosByListId(listId int) []*model.TodoStep {
	var steps []*model.TodoStep
	for _, todo := range todos {
		if todo.TodoListId == listId {
			steps = append(steps, todo)
		}
	}
	return steps
}
func GetTodoById(id int) *model.TodoStep {

	for _, step := range todos {
		if id == step.ID {
			return step
		}
	}
	return nil
}
func GetTodoListById(listId int) *model.TodoList {
	var todolist *model.TodoList
	for _, list := range todolists {
		if listId == list.ID {
			todolist = list
			break
		}
	}
	return todolist
}
func DeleteTodo(id int) {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}
	step := GetTodoById(id)
	if step != nil {
		step.DeletedAt = time.Now()
	}
}
func DeleteTodoList(id int) {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}
	list := GetTodoListById(id)
	if list != nil {
		list.DeletedAt = time.Now()
	}
}
func Complete(todoId int) {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}
	print(0)
	todo := GetTodoById(todoId)
	if todo == nil {
		return
	}
	print("1")
	todo.IsDone = true
	todo.ModifiedAt = time.Now()

	list := GetTodoListById(todo.TodoListId)
	if list == nil {
		return
	}
	list.ModifiedAt = time.Now()
	print("2")
	todos := GetTodosByListId(list.ID)
	doneCount := 0
	for _, todo := range todos {
		if todo.DeletedAt.IsZero() && todo.IsDone {
			doneCount++
		}
	}
	print("3")

	if len(todos) > 0 {
		list.Percentage = int(float64(doneCount) / float64(len(todos)) * 100)
	} else {
		list.Percentage = 0
	}
	print("4")

}

func GetTodoListsByUserId(userId int) []*model.TodoList {
	if !muLock {
		mu.Lock()
		defer MuUnlock()

	}

	user := GetUserById(userId)
	return user.TodoLists
}
