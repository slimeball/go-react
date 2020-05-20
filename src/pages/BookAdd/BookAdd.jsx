import React from "react";
import { addBookAjax, getBookbyIdAjax, updateBookAjax } from '../../service/ajax'

export default class BookAdd extends React.Component {
  componentDidMount() {
    // is edit or not
    this.isModify()
  }
  state = {
    isModify: false,
    id: '',
    bookname: '',
    author: '',
    price: '',
    sales: '',
    inventory: '',
    img_path: '',
    tip: '',
    buttonText: 'Add'
  }
  // check is modify
  isModify = () => {
    const params = new URLSearchParams(this.props.location.search);
    const type = params.get('type');
    const id = params.get('id');
    if (type !== null) {
      this.getOneBook(id);
      this.setState({
        id,
        buttonText: 'Update'
      })
      this.setState({isModify: true})
    } else {
      this.setState({
        buttonText: 'Add'
      })
      this.setState({isModify: false})
    }
  }
  // detail
  async getOneBook(id) {
    let res = await getBookbyIdAjax({ id })
    if (res.success) {
      this.showTip(res.message)
      this.setState({ ...this.state, ...res.data })
    } else {
      this.showTip(res.message)
    }
  }

  // show tip
  showTip = (msg) => {
    this.setState({ tip: msg })
    setTimeout(() => {
      this.setState({ tip: '' })
    }, 3000)
  }
  backToList = () => {
    this.props.history.push('/booklist')
  }

  // add new or update
  async addOrUpdateBook() {
    const sendData = {
      bookname: this.state.bookname,
      author: this.state.author,
      price: +this.state.price,
      sales: +this.state.sales,
      inventory: +this.state.inventory,
      img_path: this.state.img_path
    }
    let res
    if (this.isModify) {
      sendData.id = this.state.id;
      res = await updateBookAjax(sendData)
    } else {
      res = await addBookAjax(sendData)
    }
    if (res.success) {
      this.showTip(res.message)
    } else {
      this.showTip(res.message)
    }
  }
  changeEvt = (event) => {
    this.setState({
      [event.target.name]: event.target.value
    })
  }
  submitEvt = () => {
    this.addOrUpdateBook()
  }
  render() {
    return (
      <form>
        <button type="button" onClick={this.backToList}>back to list</button>
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
    <button type="button" onClick={this.submitEvt}>{this.state.buttonText}</button>
        <p>{this.state.tip}</p>
      </form>
    )
  }
}