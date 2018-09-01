import React, {Component} from 'react'
import { Letter } from './Letter';

export class Board extends Component {

    mouseDown = false;

    state = {
        lastIndex: -1,
        word: '',
        selectionMap: {}
    }

    componentDidMount() {
        this.setState({
            lastIndex: -1,
            word: '', 
            selectionMap: this.props.board.reduce((acc, _,index) => ({...acc, [index]: false}),{}),
        });
    }

    onMouseDown = () => this.mouseDown = true;
    onMouseUp = () => {
        if(this.state.word.length > 2) {
            this.props.test(this.state.word);
        }
        this.mouseDown = false;
        this.resetWord();
    };
    onBreak = () => {
        this.mouseDown = false;
        this.resetWord();
    }

    onLetterClick(index) {
        this.mouseDown = true;
        this.addLetter(index)
    }

    addLetter(index) {
        if(!this.mouseDown || this.state.lastIndex === index) {
            return;
        }

        const letter = this.props.board[index];
        if(this.lastIndex < 0){
            return this.changeWord(index, letter)
        }

        if(!!this.state.selectionMap[index]) {
            return this.resetWord();
        }

        return this.changeWord(index, letter)
    }

    resetWord() {
        this.setState({
            selectionMap: this.props.board.reduce((acc, _,index) => ({...acc, [index]: false}),{}),
            lastIndex: -1,
            word: '', 
        })
    }

    changeWord(lastIndex, letter) {
        this.setState({
            selectionMap: {...this.state.selectionMap, [lastIndex]: true},
            lastIndex,
            word: this.state.word + letter, 
        })
    }

    render() {
        return (
            <div 
                className="board" 
                onMouseDown={() => this.onMouseDown(true)} 
                onMouseUp={() => this.onMouseUp(false)} 
                onMouseLeave={() => this.onBreak()} 
            >
                {
                    this.props.board.map((letter, index) => (
                        <Letter 
                            key={index} 
                            enter={() => this.addLetter(index)} 
                            click={() => this.onLetterClick(index)}
                            letter={letter} 
                            selection={this.state.selectionMap[index]}
                            width={ 100/this.props.size }
                        ></Letter>
                    ))
                }
            </div>
        )
    }

}
