import React from "react";
export default class SignUp extends React.Component {
  state = {
    username: '',
    password: '',
    confirm_pwd: '',
    email: '',
    tip: ''
  }
  // show tip
  showTip = (msg) =>{
    this.setState({ tip: msg })
    setTimeout(()=>{
      this.setState({ tip: '' })
    }, 3000)
  }
  //  input two-way data bind
  changeEvt = (event) => {
    this.setState({
      [event.target.name]: event.target.value
    })
  }
  // link to signin
  toSignIn = () => {
    this.props.history.push('/signin')
  }
  // submit form
  submitEvt = () => {
    const sendData = {
      username: this.state.username,
      password: this.state.password,
      email: this.state.email,
    }
    fetch('http://localhost:5050/v1/signup', {
      method: 'post',
      headers: {
        'Content-Type': 'application/json'
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: JSON.stringify(sendData)
    })
      .then(res => res.json())
      .then(result => {
        if (result.success) {
          this.showTip(result.message)
        } else {
          this.showTip(result.message)
        }
      })
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
        <div>
          <label>Confirm Password:</label>
          <input type="password" name="password" value={this.state.password} onChange={this.changeEvt}></input>
        </div>
        <div>
          <label>E-mail:</label>
          <input type="text" name="email" value={this.state.email} onChange={this.changeEvt}></input>
        </div>
        <span className="signup" onClick={this.toSignIn}>SignIn</span>
        <button type="button" onClick={this.submitEvt}>SignUp</button>
        <p>{this.state.tip}</p>
      </form>
    )
  }
}