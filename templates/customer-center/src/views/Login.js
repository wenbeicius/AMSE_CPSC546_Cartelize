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

import '../assets/scss/login/login.scss';

export default function Login(props) {
    return (
        <Container className="mt-2 d-flex justify-content-center">
            <Card className="shadow border-0 mt-4 login-card">
                <CardHeader className="bg-primary text-white text-center"><h4 className="my-0">Sign In</h4></CardHeader>
                <CardBody className="pb-2">
                <Form>
                    <FormGroup>
                        <Label for="email">Email:</Label>
                        <Input type="text" name="email" placeholder="john.doe@example.com"></Input>
                    </FormGroup>
                    <FormGroup>
                        <Label for="password">Password:</Label>
                        <Input type="password" name="password" placeholder="Password"></Input>
                    </FormGroup>
                    <Button color="primary" block>Login</Button>
                    <div className="pb-0 mb-0 d-flex flex-row-reverse mt-1">
                        <a href="/register" className="text-muted small">Create Account?</a>
                        <a href="/forgot-password" className="text-muted small mr-2">Forgot Password?</a>
                    </div>
                </Form>
                </CardBody>
            </Card>
        </Container>
    )
}