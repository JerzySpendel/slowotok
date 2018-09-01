import React, { Component } from 'react';

import { Header } from './Components/Header';
import { Board } from './Components/Board';


export class Game extends Component {
    size = 4;

    state = {
        loaded: false,
        fetching: false,
        score: 0,
        board: '',
    }

    startGame(){
        if(this.state.fetching) {
            return;
        }
        this.setState((state) => ({ ...state, fetching: true }));
        fetch(`/api/board/${this.size}`)
            .then((response) => response.json())
            .then((board) => this.setState((state) => ({ ...state, board, loaded: true })))
            .finally(() => this.setState((state) => ({ ...state, fetching: false })))
    }

    testWord(word) {
        fetch(`/api/check_word/`,{method: 'POST', body: JSON.stringify({ word })})
            .then((response) => response.json())
            .then((score) => this.setState((state) => ({ ...state, score: state.score+score})))
    }

    render() {
        return (
            <section>
                <Header startGame={ () => this.startGame() }></Header>
                <div>Score: { this.state.score }</div>
                <div className="wrapper">
                    {
                        this.state.fetching || !this.state.loaded ? 
                            <div className="loader">Wait loading</div> :
                            <Board size={this.size} board={this.state.board} test={(word) => this.testWord(word)}></Board>
                    }
                </div>
            </section>
        )
    }
}