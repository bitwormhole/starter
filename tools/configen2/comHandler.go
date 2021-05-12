package configen2

import (
	"fmt"
	"go/ast"
)

type comNodeHandler interface {
	handle(n ast.Node)

	//	handle1(n *ast.StructType)    // 结构体类型
	//	handle2(n *ast.FuncType)      // 函数类型
	//	handle3(n *ast.InterfaceType) //  接口类型
}

type defaultComNodeHandler struct {
	comments *commentSet
}

func DefaultHandler(cs *commentSet) comNodeHandler {
	return &defaultComNodeHandler{comments: cs}
}

func (inst *defaultComNodeHandler) handle(n ast.Node) {

	t_gen_decl, ok := n.(*ast.GenDecl)
	if ok {
		// inst.handleComment(t_comment)
		inst.handleCommentGroup(t_gen_decl.Doc)
		return
	}

	t_comment, ok := n.(*ast.Comment)
	if ok {
		inst.handleComment(t_comment)
		return
	}

	t_comment_group, ok := n.(*ast.CommentGroup)
	if ok {
		inst.handleCommentGroup(t_comment_group)
		return
	}

	t_struct, ok := n.(*ast.StructType)
	if ok {
		inst.handleStruct(t_struct)
		return
	}

	t_interface, ok := n.(*ast.InterfaceType)
	if ok {
		inst.handleInterface(t_interface)
		return
	}

	t_func_type, ok := n.(*ast.FuncType)
	if ok {
		inst.handleFunc(t_func_type)
		return
	}

	t_func_decl, ok := n.(*ast.FuncDecl)
	if ok {
		inst.handleFuncDecl(t_func_decl)
		return
	}

	t_func_lit, ok := n.(*ast.FuncLit)
	if ok {
		inst.handleFuncLit(t_func_lit)
		return
	}

	t_import, ok := n.(*ast.ImportSpec)
	if ok {
		inst.handleImport(t_import)
		return
	}

	t_type, ok := n.(*ast.TypeSpec)
	if ok {
		inst.handleTypeSpec(t_type)
		return
	}

}

func (inst *defaultComNodeHandler) handleImport(n *ast.ImportSpec) {
	fmt.Println("[import name:", n.Name.String(), "]")
}

func (inst *defaultComNodeHandler) handleStruct(n *ast.StructType) {

	fields := n.Fields.List

	for index := range fields {
		field := fields[index]
		fmt.Println("  [field name:", field.Names, " type:", field.Type, " tag:", field.Tag, "]")
	}

	// n.Incomplete.

}

func (inst *defaultComNodeHandler) handleFunc(n *ast.FuncType) {

	fmt.Println("[func_type]")

	params := n.Params
	if params != nil {
		list := params.List
		for index := range list {
			field := list[index]
			fmt.Println("  [param name:", field.Names, " type:", field.Type, " tag:", field.Tag, "]")
		}

	}

	results := n.Results
	if results != nil {
		list := results.List
		for index := range list {
			field := list[index]
			fmt.Println("  [result name:", field.Names, " type:", field.Type, " tag:", field.Tag, "]")
		}
	}

	// innerText := inst.getInjectionProperties(n.Pos(), n.End())
	// fmt.Println("[func_inner text:", innerText, "]")

}

func (inst *defaultComNodeHandler) handleFuncDecl(n *ast.FuncDecl) {

	fmt.Println("[func_decl name:" + n.Name.Name + "]")
	body := n
	props := inst.comments.loadInjectionProperties(body.Pos(), body.End())
	fmt.Println("[comment props:", props, "]")
}

func (inst *defaultComNodeHandler) handleFuncLit(n *ast.FuncLit) {
	pos := n.Pos()
	end := n.End()
	fmt.Println("[func_lit pos:", pos, " end:", end, "]")
}

func (inst *defaultComNodeHandler) handleInterface(n *ast.InterfaceType) {

}

func (inst *defaultComNodeHandler) handleTypeSpec(n *ast.TypeSpec) {
	fmt.Println("[type name:", n.Name, " type:", n.Type, "]")
}

func (inst *defaultComNodeHandler) handleComment(n *ast.Comment) {

	if n == nil {
		return
	}

	fmt.Println("[comment text:", n.Text, "]")
}

func (inst *defaultComNodeHandler) handleCommentGroup(n *ast.CommentGroup) {

	if n == nil {
		return
	}

	fmt.Println("[commentGroup]")
	list := n.List
	if list == nil {
		return
	}
	for index := range list {
		item := list[index]
		if item == nil {
			continue
		}
		inst.handleComment(item)
	}
}
