package handlePost

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sinfor/class"
	"sinfor/db"
	"sinfor/mode"
	"sinfor/token"
	"strconv"
)

//分页查询
func ReturnPage(c *gin.Context) {
	tokenStr := c.Request.Header.Get("token")
	claims, _ := token.VerifyToken(tokenStr)
	if claims.Code[:3] == "010" {
		var pageInfor class.Msg
		err := c.ShouldBindJSON(&pageInfor)
		if err != nil {
			fmt.Println(err)
			return
		}
		if pageInfor.Message == "true" {
			num, err := strconv.Atoi(pageInfor.PageNumber)
			if err != nil {
				fmt.Println(err)
				return
			}
			rows, err := db.GetInforFromPageCode(num)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": mode.FDoSql,
					"data": struct {
					}{},
					"message": "数据库操作失败",
				})
				return
			}
			recordNum, err := db.GetPage()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": mode.FDoSql,
					"data": struct {
					}{},
					"message": "数据库操作失败",
				})
				return
			}
			infor := make([]class.PageInformation, 0, 10)
			var pageInfor class.PageInformation
			for rows.Next() == true {
				if err := rows.Scan(&pageInfor.ApplicationNum, &pageInfor.StuCode, &pageInfor.PassportWholeName, &pageInfor.ChineseName,
					&pageInfor.Specialty); err == nil {
					infor = append(infor, struct {
						ApplicationNum    string
						StuCode           string
						PassportWholeName string
						ChineseName       string
						Specialty         string
					}{ApplicationNum: pageInfor.ApplicationNum, StuCode: pageInfor.StuCode, PassportWholeName: pageInfor.PassportWholeName,
						ChineseName: pageInfor.ChineseName, Specialty: pageInfor.Specialty})
				}
			}
			var information class.ReturnPage
			information = struct {
				PageInfor    []class.PageInformation
				WholePageNum int
			}{PageInfor: infor, WholePageNum: recordNum}
			c.JSON(http.StatusOK, gin.H{
				"code":    mode.TCode,
				"data":    information,
				"message": "ok",
			})
			return
		}
		num, err := strconv.Atoi(pageInfor.PageNumber)
		statusCode, err := db.GetStatusCode(claims.Code)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"code": mode.TCode,
				"data": struct {
				}{},
				"message": err,
			})
			return
		}
		if statusCode == "false" {
			rows, err := db.GetInforFromPageCode(num)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": mode.FDoSql,
					"data": struct {
					}{},
					"message": "数据库操作失败",
				})
				return
			}
			recordNum, err := db.GetPage()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": mode.FDoSql,
					"data": struct {
					}{},
					"message": "数据库操作失败",
				})
				return
			}
			infor := make([]class.PageInformation, 0, 10)
			var pageInfor class.PageInformation
			for rows.Next() == true {
				if err := rows.Scan(&pageInfor.ApplicationNum, &pageInfor.StuCode, &pageInfor.PassportWholeName, &pageInfor.ChineseName,
					&pageInfor.Specialty); err == nil {
					infor = append(infor, struct {
						ApplicationNum    string
						StuCode           string
						PassportWholeName string
						ChineseName       string
						Specialty         string
					}{ApplicationNum: pageInfor.ApplicationNum, StuCode: pageInfor.StuCode, PassportWholeName: pageInfor.PassportWholeName,
						ChineseName: pageInfor.ChineseName, Specialty: pageInfor.Specialty})
				}
			}
			var information class.ReturnPage
			information = struct {
				PageInfor    []class.PageInformation
				WholePageNum int
			}{PageInfor: infor, WholePageNum: recordNum}
			c.JSON(http.StatusOK, gin.H{
				"code":    mode.TCode,
				"data":    information,
				"message": "ok",
			})
			return
		} else {
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": mode.FDoSql,
					"data": struct {
					}{},
					"message": "数据库操作失败",
				})
				return
			}
			rows, err := db.GetInforFromPageCode2(num, claims.Code)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": mode.FDoSql,
					"data": struct {
					}{},
					"message": "数据库操作失败",
				})
				return
			}
			recordNum, err := db.GetLocalPage()
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": mode.FDoSql,
					"data": struct {
					}{},
					"message": "数据库操作失败",
				})
				return
			}
			infor := make([]class.PageInformation, 0, 10)
			var pageInfor class.PageInformation
			for rows.Next() == true {
				if err := rows.Scan(&pageInfor.ApplicationNum, &pageInfor.StuCode, &pageInfor.PassportWholeName, &pageInfor.ChineseName,
					&pageInfor.Specialty); err == nil {
					infor = append(infor, struct {
						ApplicationNum    string
						StuCode           string
						PassportWholeName string
						ChineseName       string
						Specialty         string
					}{ApplicationNum: pageInfor.ApplicationNum, StuCode: pageInfor.StuCode, PassportWholeName: pageInfor.PassportWholeName,
						ChineseName: pageInfor.ChineseName, Specialty: pageInfor.Specialty})
				}
			}
			var information class.ReturnPage
			information = struct {
				PageInfor    []class.PageInformation
				WholePageNum int
			}{PageInfor: infor, WholePageNum: recordNum}
			c.JSON(http.StatusOK, gin.H{
				"code":    mode.TCode,
				"data":    information,
				"message": "ok",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": mode.NoneAuth,
		"data": struct {
		}{},
		"message": "该用户权限不够",
	})
	return
}
