<template>
  <div class="bg-img">
    <div class="container">
        <div class="header">
               <h2>WiseHub-How-To</h2>
        </div>
        <div class="container-fluid">
            <div id="test">
            <!-- WITH SCOPED SLOTS and INITIAL OPEN VALUES-->
                <VueFaqAccordion
                        id="accordion"
                        @categorySelect="onCategorySelect"
                        @itemSelect="onItemSelect"
                        :items="myItems"
                        v-slot="itemData"
                        initialTab="Basic Usage"
                        :initialQuestionIndex="0"
                        activeColor="#008B8B"
                        style="white-space: pre-line"
                >
                    <!-- any html -->
                    <!-- and item data if available via itemData -->
                    <!-- don't add margins to this element (animation become choppy) -->
                    <small>{{ itemData.value }}</small>
                </VueFaqAccordion>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
    import VueFaqAccordion from 'vue-faq-accordion'

    export default {
        name: "FAQ",
        components: {
            VueFaqAccordion
        },
        data () {
            return {
              myItems: [
                  {
                      title: 'What is WiseHub-Connect for and how do I use it?',
                      value: 'WiseHub-Connect is a tool for programmers, especially programming teams.\n' +
                          'It connects version management tool and continuous integration tool and gives you an all-in-one overview',
                      category: 'Basic Usage',
                  },
                  {
                      title: 'How to register?',
                      value: 'You can either register via the \'SignUp\' button or with your GitHub or GitLab account.\n' +
                          'In version v1.0 only registration via \'SignUp\' is available.\n',
                      category: 'Basic Usage'
                  },
                  {
                      title: 'What plugins are used?',
                      value: 'In version v1.0 plugins for GitHub and Drone CI are available.\n',
                      category: 'Basic Usage'
                  },
                  {
                      title: 'What credentials do I need to use these plugins?',
                      value: 'You need to have an API Token for GitHub to use WiseHub Connect.\n' +
                          'In the menu under \'Settings/Profile\' you can enter the respective credentials for GitHub and Drone CI.\n',
                      category: 'Basic Usage'
                  },
                  {
                      title: 'How to set up my own WiseHub?',
                      value: 'You clone the project from the GitHub repository onto your server first.\n' +
                          'Then you start the Application on your server via the bash script \'start_wisehub.sh\' which is located in the root directory.\n' +
                          'Please make sure that you have installed \'screen\', \'go\' and \'npm\' on your server to run WiseHub Connect.\n' +
                          'Another and probably much simpler way to do this is to run the application in a docker container.\n' +
                          'Therefore you run in root directory the following comands:\n' +
                          '\'docker-compose build\' to compile\n' +
                          '\'docker-compose up\' to run\n' +
                          '\'docker-compose down\' to stop and delete\n' +
                          'For further information please check the documentation on docker container.',
                      category: 'Admin Usage'
                  },
                  {
                      title: 'How to write your own WiseHub-Plugin?',
                      value: 'The only way to create and use your own plugin is to Host WiseHub Connect by yourself.\n' +
                          'You can find instructions on how to write a plugin in the file ./model/template/Template.go.\n',
                      category: 'Admin Usage'
                  },
              ],
            }
        },
        methods: {
            onCategorySelect (args) {
                console.log('onCategorySelect', args)
            },
            onItemSelect (args) {
                console.log('onItemSelect', args)
                this.itemData = this.myItems.value
            }
        }
    }
</script>

<style scoped lang="scss">

    .bg-img {
      /* The image used */
      background-image: url("");

      /* Control the height of the image */
      min-height: 100vh;

      /* Center and scale the image nicely */
      background-position: center;
      background-repeat: no-repeat;
      background-size: cover;
      position: relative;
    }
    .container-fluid {
      float: contour;
      horiz-align: center;
      width: 72%;

    }
    #accordion {
        width: 100%;
        padding-top: 30px;
        padding-bottom: 30px;
    }

    @media screen and (max-width: 1650px) {

    }


</style>