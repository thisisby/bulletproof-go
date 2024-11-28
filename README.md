**🚀 Bulletproof-go - The High-Performance Minimal Go Web Boilerplate**

Welcome to **Bulletproof-go** ⚡—a high-performance, minimal Go web Boilerplate crafted for building scalable, maintainable, and modern applications. LessGo empowers developers with advanced features like dynamic configuration management, inbuilt error handling, robust data validation, and Dependency Injection (DI).

Inspired by the modular architecture of frameworks like NestJs, LessGo offers a flexible structure that allows you to build your applications your way. Whether you prefer a strict controller-service-module architecture or a more fluid design, LessGo has you covered. With built-in support for CORS, CSRF, XSS protection, logging, error handling, rate limiting, caching, and job scheduling, LessGo makes it easier to develop secure, efficient, and performant web applications.

### ✨ Key Features:

1. **🔄 Concurrency & Task Management**  
   It includes a powerful task manager that allows you to run tasks in parallel or sequence, giving you full control over task management and program flow. We are committed to enhancing its robustness and extensibility.

2. **⚙️ Dynamic Configuration Management**  
   With built-in support for multiple environments (development, production, testing), it simplifies configuration management. It provides a user-friendly API for working with environment variables, offering validation and easy access to typed values like numbers and booleans.

3. **🔧 Context**  
   The Context is a powerful abstraction over the request and response objects. It simplifies API development by providing methods for handling headers, cookies, query parameters, JSON parsing, and more. Error management, HTTP redirection, and file attachments are made effortless.

4. **🧩 Flexible Dependency Injection (DI)**  
   It offers basic DI functionality that we're continuously improving. You can choose whether to bind your entire application into a single container or work with a more traditional approach, giving you flexibility in how you manage dependencies.

5. **🚀 App Initialization**  
   The App is the core of project, initializing various configurations and middleware. It exposes the main server with a simple `Listen` method to start your application.

6. **⏰ Job Scheduling**  
   LessGo supports job scheduling, leveraging the powerful Cron package to handle recurring tasks seamlessly.

7. **🔒 Comprehensive Middleware**  
   LessGo comes with essential built-in middleware, including CORS, JSON parser, cookie parser, CSRF and XSS protection, caching, file uploads, and rate limiting (both in-memory and Redis-backed).

8. **🌐 Powerful Router**  
   The built-in router handles different HTTP methods with custom handlers, supporting method chaining, sub-routers, and custom middleware.

9. **📦 Controller Interface**  
   Controllers allow you to create dedicated routers for specific endpoints, integrating seamlessly with the service layer for efficient request handling.

10. **🔍 Service Layer**  
   The Service interface is designed for encapsulating business logic. It easily binds to databases and other dependencies, streamlining the development process.

11. **🔗 Modular Design**  
   Modules bring together controllers and services, managing routes and handlers in a cohesive manner.

12. **💬 WebSocket Support (Upcoming)**  
   We are actively working on integrating WebSocket support to make real-time communication in your applications even easier.


### 🙌 Acknowledgments

Initial idea was borrowed from this guy: https://github.com/hokamsingh 

### 📜 License

This project is licensed under the MIT License - see the LICENSE file for details.

**Happy coding!** 🎉

--- 
