import React, { Component } from "react";
import { Container, Button, Form, Col, Row } from "react-bootstrap";

const handleSubmit = (event) => {
  event.preventDefault(); // Prevents the default form submission

  // Custom logic: You can perform form validation or data handling here

  // Example: Log form data
  const formData = new FormData(event.target);
  for (let [name, value] of formData.entries()) {
    console.log(`${name}: ${value}`);
  }
};
class Register extends Component {
  render() {
    return (
      <Container>
        <h1>Register Here</h1>
        <Row className="justify-content-center">
          <Col md={6}>
            <Form onSubmit={handleSubmit}>
              <Form.Group controlId="firstName" className="mb-3">
                <Form.Label>First Name</Form.Label>
                <Form.Control type="text" placeholder="Enter your first name" />
              </Form.Group>
              <Form.Group controlId="lastName" className="mb-3">
                <Form.Label>Last Name</Form.Label>
                <Form.Control type="text" placeholder="Enter your last name" />
              </Form.Group>
              <Form.Group controlId="username" className="mb-3">
                <Form.Label>Username</Form.Label>
                <Form.Control type="text" placeholder="Enter your username" />
              </Form.Group>
              <Form.Group controlId="password" className="mb-3">
                <Form.Label>Password</Form.Label>
                <Form.Control type="password" placeholder="Password" />
              </Form.Group>
              <Form.Group controlId="age" className="mb-3">
                <Form.Label>Your Age</Form.Label>
                <Form.Control type="number" placeholder="Your age?" />
              </Form.Group>
              <Button variant="primary" type="submit">
                Submit
              </Button>
            </Form>
          </Col>
        </Row>
      </Container>
    )
  }
}

export default Register;