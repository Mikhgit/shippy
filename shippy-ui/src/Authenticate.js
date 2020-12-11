// shippy-ui/src/Authenticate.js
import React from 'react';

class Authenticate extends React.Component {

    constructor(props) {
        super(props);
    }

    state = {
        authenticated: false,
        email: '',
        password: '',
        err: '',
    }

    login = () => {
        fetch(`http://localhost:8080/auth/auth`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                //     request: {
                email: this.state.email,
                password: this.state.password,
                // },
                // service: 'go.micro.api.auth',
                // method: 'Auth.Auth',
            }),
        })
            .then(res => res.json())
            .then(res => {
                this.props.onAuth(res.token);
                this.setState({
                    token: res.token,
                    authenticated: true,
                });
                localStorage.setItem('token', res.token);
            })
            .catch(err => this.setState({err, authenticated: false,}));
    }

    signup = () => {
        fetch(`http://localhost:8080/auth/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                // request: {
                email: this.state.email,
                password: this.state.password,
                name: this.state.name,
                // },
                // method: 'Auth.Create',
                // service: 'go.micro.api.auth',
            }),
        })
            .then((res) => res.json())
            .then((res) => {
                this.props.onAuth(res.token);
                this.setState({
                    token: res.token,
                    authenticated: false,
                });
            })
            .catch(err => this.setState({err, authenticated: false,}));
    }

    setEmail = e => {
        this.setState({
            email: e.target.value,
        });
    }

    setPassword = e => {
        this.setState({
            password: e.target.value,
        });
    }

    setName = e => {
        this.setState({
            name: e.target.value,
        });
    }

    render() {
        return (
            <div className='Authenticate'>
                <div className='Login'>
                    <div className='form-group'>
                        <input
                            type="email"
                            onChange={this.setEmail}
                            placeholder='E-Mail'
                            className='form-control'/>
                    </div>
                    <div className='form-group'>
                        <input
                            type="password"
                            onChange={this.setPassword}
                            placeholder='Password'
                            className='form-control'/>
                    </div>
                    <button className='btn btn-primary' onClick={this.login}>Login</button>
                    <br/><br/>
                </div>
                <div className='Sign-up'>
                    <div className='form-group'>
                        <input
                            type='input'
                            onChange={this.setName}
                            placeholder='Name'
                            className='form-control'/>
                    </div>
                    <div className='form-group'>
                        <input
                            type='email'
                            onChange={this.setEmail}
                            placeholder='E-Mail'
                            className='form-control'/>
                    </div>
                    <div className='form-group'>
                        <input
                            type='password'
                            onChange={this.setPassword}
                            placeholder='Password'
                            className='form-control'/>
                    </div>
                    <button className='btn btn-primary' onClick={this.signup}>Sign-up</button>
                </div>
            </div>
        );
    }
}

export default Authenticate;