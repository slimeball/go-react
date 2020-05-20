import React from "react";
import './BookList.scss';
import { getBookListAjax, deleteBookAjax } from '../../service/ajax'

export default class BookList extends React.Component {
  componentDidMount() {
    this.getList()
  }
  state = {
    booklist: [],
    tip: ''
  }
  // show tip
  showTip = (msg) =>{
    this.setState({ tip: msg })
    setTimeout(()=>{
      this.setState({ tip: '' })
    }, 3000)
  }
  // add new book
  toAddBook = () => {
    this.props.history.push('/bookadd')
  }
  async getList() {
    let res = await getBookListAjax()
    if (res.success) {
      this.setState({ booklist: res.data })
      this.showTip(res.message)
    } else {
      this.showTip(res.message)
    }
  }
  async deletebook(id){
    let res = await deleteBookAjax(id)
    if (res.success) {
      this.getList()
      this.showTip(res.message)
    } else {
      this.showTip(res.message)
    }
  }
  toEdit = (id) =>{
    this.props.history.push('/bookadd?type=edit&id='+id)
  }
  // delete
  toDelete = (id) => {
    if(window.confirm("Do you really want to leave?")){
      let sendData = {
        id:id
      }
      this.deletebook(sendData)
    }
  }
  render() {
    return (
      <div className="book-list">
        <button className="weui-btn weui-btn_mini weui-btn_default" onClick={this.toAddBook}>Add book</button>
        <p className="tip">{this.state.tip}</p>
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Bookname</th>
              <th>Author</th>
              <th>Price</th>
              <th>Sales</th>
              <th>Inventory</th>
              <th>ImgPath</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {this.state.booklist.map(item => {
            return <tr key={item.id}>
              <td>{item.id}</td>
              <td>{item.bookname}</td>
              <td>{item.author}</td>
              <td>{item.price}</td>
              <td>{item.sales}</td>
              <td>{item.inventory}</td>
              <td>{item.img_path}</td>
              <td><button type="button" className="weui-btn weui-btn_mini weui-btn_primary" onClick={this.toEdit.bind(this,item.id)}>Edit</button><button type="button" className="weui-btn weui-btn_mini weui-btn_warn" onClick={this.toDelete.bind(this, item.id)}>Delete</button></td>
            </tr>})
          }
          </tbody>
        </table>
      </div>
    )
  }
}