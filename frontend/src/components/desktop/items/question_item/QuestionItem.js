import './QuestionItem.css';
import React from 'react';

export function QuestionItem({title, question}) {
    return (
        <div className='question-item'>
            <div className='question-item-description-wrapper'>
                <h1 className='question-item-description-title'>Заголовок: {title}</h1>
                <p className='question-item-description-sub'> Текст сообщения: {question}</p>
            </div>
        </div>
    )
}