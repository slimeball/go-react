import React from "react";
import './BookList.scss';
import { getBookListAjax, deleteBookAjax } from '../../service/ajax'

export default class BookList extends React.Component {
  componentDidMount() {
    this.getList()
  }
  state = {
    pageno: 1,
    pagesize: 5,
    pagetotal: 0,
    pagetotalnum: 0,
    pageArr: [],
    bookRes: [],
    tip: ''
  }
  // show tip
  showTip = (msg) => {
    this.setState({ tip: msg })
    setTimeout(() => {
      this.setState({ tip: '' })
    }, 3000)
  }
  // add new book
  toAddBook = () => {
    this.props.history.push('/bookadd')
  }
  async getList() {
    let sendData = {
      pageno: this.state.pageno,
      pagesize: this.state.pagesize
    }
    let res = await getBookListAjax(sendData)
    if (res.success) {
      let page = res.data.pagetotalnum
      let arr = [];
      for (let i = 1; i <= page; i++) {
        arr.push(i)
      }
      this.setState({
        bookRes: res.data.book,
        pagetotalnum: res.data.pagetotalnum,
        pagetotal: res.data.pagetotal,
        pageArr: arr
      })
      this.showTip(res.message)
    } else {
      this.showTip(res.message)
    }
  }
  async deletebook(id) {
    let res = await deleteBookAjax(id)
    if (res.success) {
      this.getList()
      this.showTip(res.message)
    } else {
      this.showTip(res.message)
    }
  }
  // change page
  async pageChange (no) {
    await this.setState({
      pageno: no
    })
    this.getList()
  }

  toEdit = (id) => {
    this.props.history.push('/bookadd?type=edit&id=' + id)
  }
  // delete
  toDelete = (id) => {
    if (window.confirm("Do you really want to leave?")) {
      let sendData = {
        id: id
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
            {this.state.bookRes.map(item => {
              return <tr key={item.id}>
                <td>{item.id}</td>
                <td>{item.bookname}</td>
                <td>{item.author}</td>
                <td>{item.price}</td>
                <td>{item.sales}</td>
                <td>{item.inventory}</td>
                <td>{item.img_path}</td>
                <td><button type="button" className="weui-btn weui-btn_mini weui-btn_primary" onClick={this.toEdit.bind(this, item.id)}>Edit</button><button type="button" className="weui-btn weui-btn_mini weui-btn_warn" onClick={this.toDelete.bind(this, item.id)}>Delete</button></td>
              </tr>
            })
            }
          </tbody>
        </table>
        <div className="paginationContainer">
          {this.state.pageArr.map(item => {
            return <span className={this.state.pageno === item ? 'pagination cur_no': 'pagination'} key={item} onClick={this.pageChange.bind(this, item)}>{item}</span>
          })
          }
          <div>current page<span>{this.state.pageno}</span></div>
          <div>total page<span>{this.state.pagetotalnum}</span></div>
          <div>total result<span>{this.state.pagetotal}</span></div>
        </div>
      </div>
    )
  }
}