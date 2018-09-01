import React, { Component } from 'react';

import {Header} from './Components/Header';
import {Board} from './Components/Board';


export class Game extends Component {
    size = 4;

    state = {
        fetching: false,
        board: '',
    }

    startGame(){
        if(this.state.fetching) {
            return;
        }
        this.setState((state) => ({ ...state, fetching: true }));
        fetch(`/board/${this.size}/`)
            .then((response) => response.json())
            .then((board) => this.setState((state) => ({ ...state, board })))
            .finally(() => this.setState((state) => ({ ...state, fetching: false })))
    }

    render() {
        return (
            <section>
                <Header startGame={() => this.startGame() }></Header>
                {this.state.fetching ? <div className="loader">Wait loading</div> : <Board board={this.state.board}></Board>}
            </section>
        )
    }
}