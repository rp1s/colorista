
# Colorista

Небольшая Go-библиотека для работы с цветами и градиентами в терминале.

## Установка

Собрать локально:

```bash
go build ./...
```

Установить бинарь (если проект содержит команду `main`):

```bash
go install ./...
```

## Быстрый пример использования

Пример кода — создание градиента и вывод значений в терминал:

```go
package main

import (
	"fmt"
	"github.com/rp1s/colorista"
)

func main() {
	c1 := colorista.NewColor(255, 0, 0) // красный
	c2 := colorista.NewColor(0, 0, 255) // синий

	g := colorista.NewGradient(c1, c2)

	// Вывести несколько точек градиента
	for i := 0; i <= 10; i++ {
		t := float64(i) / 10.0
		col := g.At(t)
		fmt.Println(col.String())
	}
}
```

Замените `github.com/your/module/path` на ваш реальный модуль из `go.mod`.

## API (кратко)

- `NewColor(r, g, b int)` - создаёт цвет.
- `Gradient(text string, stops []GradientPos, styles ...Style)` - создание строки гридиента букв
- `BgGradient(text string, stops []GradientPos, styles ...Style)` - создание строки гридиента фона
- `FullGradient(text string, fg []GradientPos, bg []GradientPos, styles ...Style)` - создание строки гридиента букв и фона


## Контакты и вклад

Если хотите внести изменения - форкните репозиторий и пришлите pull request.

