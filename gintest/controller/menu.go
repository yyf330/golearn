package controller

import (
	"net/http"
	"strconv"

	"test/gintest/db"
	"test/gintest/models"

	"github.com/gin-gonic/gin"
)

// MenuController handle menu request
type MenuController struct {
	MenuDao db.MenuDao
}

// Index handle /menu
func (MenuController) Index(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/menu/menu.html", gin.H{})
}

// List query all menu
func (mc MenuController) List(c *gin.Context) {
	var page models.Page
	if err := c.Bind(&page); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	menus, err := mc.MenuDao.List(page)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, menus)
}

// Remove delete menu
func (mc MenuController) Remove(c *gin.Context) {
	menuID := c.PostForm("menuId")
	if menuID == "" {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": "参数错误",
		})
		return
	}
	id, err := strconv.ParseInt(menuID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = mc.MenuDao.Delete(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "success",
	})
}

// ToAdd to add menu page
func (MenuController) ToAdd(c *gin.Context) {
	r.HTML(c.Writer, http.StatusOK, "system/menu/menu_add.html", gin.H{})
}

// SelectMenuTreeList query menu
func (mc MenuController) SelectMenuTreeList(c *gin.Context) {
	menus, err := mc.MenuDao.SelectMenuTreeList()
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	treeNode := models.ZTreeNode{
		ID:   0,
		Name: "顶级",
		Open: true,
		Pid:  0,
	}
	menus = append(menus, treeNode)
	r.JSON(c.Writer, http.StatusOK, menus)
}

// ToEdit update menu
func (mc MenuController) ToEdit(c *gin.Context) {
	menuID := c.Param("menuId")
	id, err := strconv.ParseInt(menuID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	menu, err := mc.MenuDao.Get(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if menu.Pcode != "0" {
		pMenu, err := mc.MenuDao.GetByPcode(menu.Pcode)
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		menu.PcodeName = pMenu.Name
		menu.Pid = pMenu.Id
	}

	r.HTML(c.Writer, http.StatusOK, "system/menu/menu_edit.html", gin.H{
		"menu": menu,
	})
}

// Add add a menu
func (mc MenuController) Add(c *gin.Context) {
	var menu models.Menu
	if err := c.Bind(&menu); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := menuSetPcode(&menu)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = mc.MenuDao.Save(menu)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "success",
	})
}

// Edit update menu
func (mc MenuController) Edit(c *gin.Context) {
	var menu models.Menu
	if err := c.Bind(&menu); err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := menuSetPcode(&menu)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = mc.MenuDao.Update(menu)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "success",
	})
}

// TreeListByRoleID query menu by roleid
func (mc MenuController) TreeListByRoleID(c *gin.Context) {
	roleID := c.Param("roleId")
	id, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		r.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	menuIDs, err := mc.MenuDao.GetMenuIdsByRoleID(id)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if len(menuIDs) == 0 {
		nodes, err := mc.MenuDao.SelectMenuTreeList()
		if err != nil {
			r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		r.JSON(c.Writer, http.StatusOK, nodes)
		return
	}
	nodes, err := mc.MenuDao.GetMenusByMenuIDs(menuIDs)
	if err != nil {
		r.JSON(c.Writer, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(c.Writer, http.StatusOK, nodes)
}

func menuSetPcode(menu *models.Menu) error {
	if menu.Pcode == "" || menu.Pcode == "0" {
		menu.Pcode = "0"
		menu.Pcodes = "[0],"
		menu.Levels = 1
	}

	pid, err := strconv.ParseInt(menu.Pcode, 10, 64)
	if err != nil {
		return err
	}
	var menuDao db.MenuDao
	pMenu, err := menuDao.Get(pid)
	if err != nil {
		return err
	}
	if menu.Code == pMenu.Code {
		return err
	}
	menu.Pcode = pMenu.Code
	menu.Levels = pMenu.Levels + 1
	menu.Pcodes = pMenu.Pcodes + "[" + pMenu.Code + "],"
	menu.Status = 1
	return nil
}
