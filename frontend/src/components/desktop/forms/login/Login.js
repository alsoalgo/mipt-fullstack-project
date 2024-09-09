import './Login.css';
import React, { useEffect } from 'react';
import { useState } from 'react';
import { Logo } from './logo/Logo';
import { Link, useNavigate } from 'react-router-dom';
import { ApiQuery, TokenCheck } from '../../../../services/Api';

export function Form() {
    let navigate = useNavigate();
    
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    useEffect(() => {
        TokenCheck().then((ok) => {
            if (ok) {
                navigate('/');
            }
        });
    }, [])

    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    }

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    }
    
    const submitForm = async (event) => {
        event.preventDefault()
        
        ApiQuery('/login', {
            method: 'post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                'email': email,
                'password': password,
            })
        }).then((resp) => {
            if (resp.status == "failed" || resp.statusCode != 200) {
                var total = resp.statusCode + " " + resp.message;
                alert(total);
                return;
            }

            navigate('/');
        })
        
    }

    return (
        <div className='forms-login'>
            <div className='forms-login-header'>
                <Logo />
                <div className='forms-login-header-invitation'>
                    Войдите или <Link to='/registration' className='forms-login-header-invitation-link'>зарегистрируйтесь</Link>
                </div>
            </div>
            <div className='forms-login-inputs'>
                <input key="forms-login-input-email" type='email' className='forms-login-input' placeholder='email' onChange={handleEmailChange}></input>
                <input key="forms-login-input-password" type='password' className='forms-login-input' placeholder='password' onChange={handlePasswordChange}></input>
            </div>
            <button className='forms-login-submit' onClick={submitForm}>Войти</button>
        </div>
    );
}