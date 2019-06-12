import React from 'react';

import {
    Card,
    CardBody,
    Container,
    Form,
    FormGroup,
    Input,
    Label,
    CardHeader,
    Button
} from 'reactstrap';
export default function Register(props) {
    return (
        <Container className="mt-2 d-flex justify-content-center">
            <Card className="shadow border-0 mt-4 login-card">
                <CardHeader className="bg-primary text-white text-center"><h4 className="my-0">Create a new account</h4></CardHeader>
                <CardBody className="pb-2">
                <Form>
                    <FormGroup>
                        <Label for="email">Your Email:</Label>
                        <Input type="text" name="email" placeholder="john.doe@example.com"></Input>
                    </FormGroup>
                    <FormGroup>
                        <Label for="name">Enter your name</Label>
                        <Input type="text" name="name" placeholder="John Doe"></Input>
                    </FormGroup>
                    <FormGroup>
                        <Label for="password">Create a password:</Label>
                        <Input type="password" name="password" placeholder = "Password!"></Input>
                    </FormGroup>
                    <FormGroup>
                        <Label for="confPassword">Confirm your password:</Label>
                        <Input type="password" name="confPassword" placeholder = "Password!"></Input>
                    </FormGroup>
                    <Button color="primary" block>Create Account!</Button>
                    <div className="pb-0 mb-0 d-flex flex-row-reverse mt-1">
                        <a href="/login" className="text-muted small">Already have a account?</a>
                    </div>
                </Form>
                </CardBody>
            </Card>
        </Container>
    )
}
