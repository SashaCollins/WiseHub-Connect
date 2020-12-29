### Continuous Integration

https://ieeexplore.ieee.org/abstract/document/6802994
29.12.2020 03:23

Continuous integration (CI) is the practice of automating the integration of code changes from multiple contributors into a single software project. It’s a primary DevOps best practice, allowing developers to frequently merge code changes into a central repository where builds and tests then run. Automated tools are used to assert the new code’s correctness before integration.

A source code version control system is the crux of the CI process. The version control system is also supplemented with other checks like automated code quality tests, syntax style review tools, and more.  

CI helps to scale up headcount and delivery output of engineering teams. Introducing CI to the aforementioned scenario allows software developers to work independently on features in parallel. When they are ready to merge these features into the end product, they can do so independently and rapidly. CI is a valuable and well-established practice in modern, high performance software engineering organizations.

CI is generally used alongside an agile software development workflow. An organization will compile list of tasks that comprise a product roadmap. These tasks are then distributed amongst software engineering team members for delivery. Using CI enables these software development tasks to be developed independently and in parallel amongst the assigned developers. Once one of theses tasks is complete, a developer will introduce that new work to the CI system to be integrated with the rest of the project.

CI vs Continuous Deployment vs Continuous Delivery (Foto)
Continuous integration, deployment, and delivery are three phases of an automated software release pipeline, including a DevOps pipeline. These three phases take software from idea to delivery to the end-user. The integration phase is the first step in the process. Continuous integration covers the process of multiple developers attempting to merge their code changes with the master code repository of a project.

Continuous delivery is the next extension of continuous integration. The delivery phase is responsible for packaging an artifact together to be delivered to end-users. This phase runs automated building tools to generate this artifact. This build phase is kept ‘green,’ which means that the artifact should be ready to deploy to users at any given time.

Continuous deployment is the final phase of the pipeline. The deployment phase is responsible for automatically launching and distributing the software artifact to end-users. At deployment time, the artifact has successfully passed the integration and delivery phases. Now it is time to automatically deploy or distribute the artifact. This will happen through scripts or tools that automatically move the artifact to public servers or to another mechanism of distribution, like an app store.

https://www.atlassian.com/continuous-delivery/continuous-integration
29.12.2020 03:32

https://www.redhat.com/de/topics/devops/what-is-ci-cd
