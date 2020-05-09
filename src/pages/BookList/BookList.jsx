import React from "react";

export default class BookList extends React.Component {
  toAddBook = () => {
    this.props.history.push('/bookadd')
  }
  render() {
    return (
      <div>
        <button onClick={this.toAddBook}>Add book</button>
        <ul>
          <li></li>
        </ul>
      </div>
    )
  }
}