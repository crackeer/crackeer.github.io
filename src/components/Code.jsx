import React from 'react';
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter';
class Code extends React.Component {
    render() {
        return <SyntaxHighlighter showLineNumbers={true}
            startingLineNumber={1}
            language={this.props.lang}
            lineNumberStyle={{ color: '#ddd' }}
            wrapLines={true}
        >
            {this.props.children.replace(/^\s+|\s+$/g, '')}
        </SyntaxHighlighter>
    }
}

export default Code