package main

import ("github.com/gin-gonic/gin"
        "gopkg.in/gorp.v1"
        "github.com/go-sql-driver/mysql"
)

func main() {
  router:=gin.Default()
  // dbMap:=initDB()
  router.GET("/login", func (c *gin.Context)  {
    c.String(200,"okz")
  })
  router.Run(":2000")
}

// func initDb() *Db  {
//
// }
