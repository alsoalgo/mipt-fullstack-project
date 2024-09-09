import "./Profile.css";
import React, { useEffect, useState } from "react";
import { Header } from "../../header/Header";
import { Form } from "../../forms/profile/Profile";
import { OrderItem } from "../../items/order_item/OrderItem";
import { QuestionItem } from "../../items/question_item/QuestionItem";
import { ApiQuery, TokenCheck } from '../../../../services/Api';
import { useNavigate } from "react-router-dom";

function QueryOrders(set) {
    ApiQuery('/orders', {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
        },
    }).then((resp) => {
        if (resp.status != "failed") {
            var orders = resp.data.orders;
            set(orders.map(
                (item, index) => {
                    var hotel = item.hotel;
                    return <OrderItem 
                        key={index + 1} 
                        title={hotel.title} 
                        description={hotel.description} 
                        img={hotel.imageUrl} 
                    />
                }
                
            ));
        }
    });
}

function QueryQuestions(set) {
    ApiQuery('/questions', {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
        },
    }).then((resp) => {
        if (resp.status != "failed") {
            var questions = resp.data.questions;
            set(questions.map(
                (item, index) => {
                    var question = item;
                    return <QuestionItem 
                            key={index + 1} 
                            title={question.title}
                            question={question.question}
                        />;
                }
                
            ));
        }
    });
}

export function Profile() {
    let navigate = useNavigate();
    const [orders, setOrders] = useState([]);
    const [questions, setQuestions] = useState([]);

    useEffect(() => {
        TokenCheck().then((ok) => {
            if (!ok) {
                navigate('/login');
            }
        });

        QueryOrders(setOrders);
        QueryQuestions(setQuestions);
    }, []);

    return (
        <div className="profile-page"> 
            <Header />
            <div className="profile-wrapper">
                <div className="profile-order-question-wrapper">
                    {   orders.length > 0 && 
                        <>
                            <h1> Ваши бронирования </h1>
                            <div className="profile-orders">
                                { orders }
                            </div>
                        </> 
                    }
                    {   questions.length > 0 && 
                        <>
                            <h1> Ваши обращения </h1>
                            <div className="profile-questions">
                                { questions }
                            </div>
                        </> 
                    }
                </div>
                <div className="profile-form-wrapper">
                    <h1> Информация </h1>
                    <Form/>
                </div>
            </div>
        </div>
    );
}