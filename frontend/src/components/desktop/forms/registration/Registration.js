import './Registration.css';
import React from 'react';
import { useState } from 'react';
import { Logo } from './logo/Logo';
import { ApiQuery } from '../../../../services/Api';
import { useNavigate } from 'react-router-dom';

export function Form() {
    let navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [againPassword, setAgainPassword] = useState('');


    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    }

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    }

    const handleAgainPasswordChange = (event) => {
        setAgainPassword(event.target.value);
    }

    const submitForm = async (event) => {
        if (password != againPassword) {
            alert('Password is not correct!');
            return;
        }

        ApiQuery('/register', {
            method: 'post',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                "email": email,
                "password": password,
            })
        }).then((resp) => {
            if (resp.statusCode != 200 || resp.status == "failed") {
                var total = resp.statusCode + " " + resp.message;
                alert(total);
                return;
            }

            navigate('/login');
        })
    }

    return (
        <div className='forms-registration'>
            <div className='forms-registration-header'>
                <Logo />
                <div className='forms-registration-header-invitation'>
                    Регистрация нового аккаунта
                </div>
            </div>
            <div className='forms-registration-inputs'>
                <input key="forms-registration-input-email" type='email' className='forms-registration-input' placeholder='email' onChange={handleEmailChange}></input>
                <input key="forms-registration-input-password" type='password' className='forms-registration-input' placeholder='password' onChange={handlePasswordChange}></input>
                <input key="forms-registration-input-again-password" type='password' className='forms-registration-input' placeholder='again password' onChange={handleAgainPasswordChange}></input>
            </div>
            <button className='forms-registration-submit' onClick={submitForm}>Зарегистрироваться</button>
        </div>
    );
}