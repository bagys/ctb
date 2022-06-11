## Usage

```
go get github.com/bagys/ctb
```

```go
import (
    "fmt"
    "math/rand"
    "time"
    "github.com/bagys/ctb"
    "github.com/gookit/color"
)

func main(){
    colors := []color.Color{color.Green, color.BgYellow}
    rand.Seed(time.Now().Unix())

    t := ctb.NewTable()
    t.SetPrefixDisable(true)

    t.SetTab([]ctb.LineData{
        {Data: "Host"},
        {Data: "IP"},
        {Data: "STATUS"},
    })

    for i := 0; i < 10; i++ {
        t.SetDataOne([]ctb.LineData{
            {Data: fmt.Sprintf("a-%d", i), Color: colors[rand.Intn(2)]},
            {Data: fmt.Sprintf("b-%d", i)},
            {Data: fmt.Sprintf("c-%d", i)},
        })
    }
    t.Print()
}
```



```go
import (
    "github.com/bagys/ctb"
    "github.com/gookit/color"
)

func main(){
    tab := []ctb.LineData{
        {Data: "Host"},
        {Data: "IP"},
        {Data: "STATUS"},
    }
    var dataAll [][]ctb.LineData
    for i := 0; i < 10; i++ {
        dataAll = append(dataAll, []ctb.LineData{
            {Data: "AA"},
            {Data: "BB"},
            {Data: "CC"},
        })
    }

    dataOne := []ctb.LineData{
        {Data: "XXX"},
        {Data: "XXX", Color: color.Red},
        {Data: "XXX"},
    }

    ctb.NewTable(
        ctb.WithDataOne(dataOne),
        ctb.WithDataAll(dataAll),
        ctb.WithDataOne(dataOne),
        ctb.WithTab(tab),
    ).Print()
}
```

