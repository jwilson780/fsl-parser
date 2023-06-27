grammar FSL;

program: statement*;

statement: variableDeclaration
         | functionDefinition
         | command;

variableDeclaration: ID ':' value;

functionDefinition: ID ':' '[' command* ']';

command: 'cmd' ':' (functionCall | arithmeticOperation | crudOperation | printOperation);

arithmeticOperation: (addOperation | subtractOperation | multiplyOperation | divideOperation);

crudOperation: (createOperation | updateOperation | deleteOperation);


addOperation: 'add' ',' 'id' ':' identificationParameter ',' 'operand1' ':' operationParameter ',' 'operand2' ':' operationParameter;

subtractOperation: 'subtract' ',' 'id' ':' identificationParameter ',' 'operand1' ':' operationParameter ',' 'operand2' ':' operationParameter;

multiplyOperation: 'multiply' ',' 'id' ':' identificationParameter ',' 'operand1' ':' operationParameter ',' 'operand2' ':' operationParameter;

divideOperation: 'divide' ',' 'id' ':' identificationParameter ',' 'operand1' ':' operationParameter ',' 'operand2' ':' operationParameter;


createOperation: 'create' ',' 'id' ':' ID ',' 'value' ':' value;

updateOperation: 'update' ',' 'id' ':' ID ',' 'value' ':' (value | variableReference);

deleteOperation: 'delete' ',' 'id' ':' ID;


printOperation: 'print' ',' 'value' ':' (value | variableReference);


functionCall: '#' ID (',' ('id' | ID) ':' passedParameter)*;

passedParameter: (ID | variableReference);

// functions can accept references to variables, actual values, or parameters
operationParameter: (variableReference | value | functionParameter);

identificationParameter: (ID | functionParameter);

functionParameter: ('$' ('id' | ID));

variableReference: '#' ID;

value: INT | FLOAT;

ID: [a-zA-Z_][a-zA-Z_0-9]*;
INT: [0-9]+;
FLOAT: [0-9]+ '.' [0-9]+;
WS: [ \t\r\n]+ -> skip;
