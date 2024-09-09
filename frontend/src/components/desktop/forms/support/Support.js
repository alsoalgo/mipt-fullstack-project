import './Support.css';
import React from 'react';
import { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { ApiQuery } from '../../../../services/Api';

function SendForm(title, question) {
    ApiQuery('/question', {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "title" : title,
            "question": question,
        })
    }
    ).then((resp) => {});
}

export function Form() {
    let navigate = useNavigate();
    const [title, setTitle] = useState('');
    const [appeal, setAppeal] = useState('');

    const handleTitleChange = (event) => {
        setTitle(event.target.value);
    }

    const handleAppealChange = (event) => {
        setAppeal(event.target.value);
    }

    const submitForm = async (event) => {
        SendForm(title, appeal);
        navigate('/profile');
    }

    return (
        <div className='forms-support'>
            <div className='forms-support-header'>
                <h1>У вас что-то случилось?</h1>
                <h2>Расскажите нам об этом!</h2>
            </div>
            <div className='forms-support-inputs'>
                <input type='text' className='forms-support-title-input' placeholder='Тема сообщения' onChange={handleTitleChange} value={title}></input>
                <textarea type='text' className='forms-support-appeal-input' placeholder='Введите сообщение' onChange={handleAppealChange} value={appeal}></textarea>
            </div>
            <div className="forms-support-submit-wrapper">
                <button className='forms-support-submit' onClick={submitForm}>Отправить</button>
            </div>
        </div>
    );
}