import React from "react";
import './SignIn.scss';
import { signinAjax } from '../../service/ajax'
export default class SignIn extends React.Component {
  state = {
    username: '',
    password: '',
    tip: ''
  }
  // show tip
  showTip = (msg) => {
    this.setState({ tip: msg })
    setTimeout(() => {
      this.setState({ tip: '' })
    }, 3000)
  }
  //  input two-way data bind
  changeEvt = (event) => {
    this.setState({
      [event.target.name]: event.target.value
    })
  }
  // link to signup
  toSignUp = () => {
    this.props.history.push('/signup')
  }
  // submit form
  async doSignin() {
    const sendData = {
      username: this.state.username,
      password: this.state.password
    }
    let res = await signinAjax(sendData)
    if (res.success) {
      sessionStorage.setItem('signin', true)
      this.props.history.push('/booklist')
    } else {
      this.showTip(res.message)
    }
  }
  submitEvt = () => {
    this.doSignin()
    // fetch('/signin', {
    //   method: 'post',
    //   headers: {
    //     'Content-Type': 'application/json'
    //     // 'Content-Type': 'application/x-www-form-urlencoded',
    //   },
    //   body: JSON.stringify(sendData)
    // })
    //   .then(res => res.json())
    //   .then(result => {
    //     if (result.success) {
    //       sessionStorage.setItem('signin', true)
    //       this.props.history.push('/booklist')
    //     } else {
    //       this.showTip(result.message)
    //     }
    //   })
  }

  render() {
    return (
      <form>
        <div>
          <label>Username:</label>
          <input type="text" name="username" value={this.state.username} onChange={this.changeEvt}></input>
        </div>
        <div>
          <label>Password:</label>
          <input type="password" name="password" value={this.state.password} onChange={this.changeEvt}></input>
        </div>
        <span className="signup" onClick={this.toSignUp}>SignUp</span>
        <button type="button" onClick={this.submitEvt}>SignIn</button>
        <p>{this.state.tip}</p>
      </form>
    )
  }
}