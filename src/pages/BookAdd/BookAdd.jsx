import React from "react";
import { addBookAjax } from '../../service/ajax'

export default class BookAdd extends React.Component {
  state = {
    bookname: '',
    author: '',
    price: '',
    sales: '',
    inventory: '',
    img_path: '',
    tip: ''
  }
  async addBook() {
    const sendData = {
      bookname: this.state.bookname,
      author: this.state.author,
      price: +this.state.price,
      sales: +this.state.sales,
      inventory: +this.state.inventory,
      img_path: this.state.img_path
    }
    console.log(sendData)
    let res = await addBookAjax(sendData)
    console.log(res)
  }
  changeEvt = (event) => {
    this.setState({
      [event.target.name]: event.target.value
    })
  }
  submitEvt = () => {
    this.addBook()
  }
  render() {
    return (
      <form>
        <div>
          <label>Bookname:</label>
          <input type="text" name="bookname" value={this.state.bookname} onChange={this.changeEvt}></input>
        </div>
        <div>
          <label>Author:</label>
          <input type="text" name="author" value={this.state.author} onChange={this.changeEvt}></input>
        </div>
        <div>
          <label>Price:</label>
          <input type="text" name="price" value={this.state.price} onChange={this.changeEvt}></input>
        </div>
        <div>
          <label>Sales:</label>
          <input type="text" name="sales" value={this.state.sales} onChange={this.changeEvt}></input>
        </div>
        <div>
          <label>Inventory:</label>
          <input type="text" name="inventory" value={this.state.inventory} onChange={this.changeEvt}></input>
        </div>
        <div>
          <label>Photo:</label>
          <input type="text" name="img_path" value={this.state.img_path} onChange={this.changeEvt}></input>
        </div>
        <button type="button" onClick={this.submitEvt}>Add</button>
        <p>{this.state.tips}</p>
      </form>
    )
  }
}