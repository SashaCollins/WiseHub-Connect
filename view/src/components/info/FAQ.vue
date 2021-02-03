<template>
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
                >
                    <!-- any html -->
                    <!-- and item data if available via itemData -->
                    <!-- don't add margins to this element (animation become choppy) -->
                    <small>{{ itemData.value }}</small>
                </VueFaqAccordion>
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
                        value: 'Given a 24-hour day and 360 degrees of longitude around the Earth... biowefugwrpeo fugperug54745475perugerugp0erjhfds giuesdhgi uewhrjgoieiugh iuregheiorgji ewrhgoie;rhgi roightoir jgioe',
                        category: 'Basic Usage',
                    },
                    {
                        title: 'How to register?',
                        value: 'You can either register via the \'SignUp\' button or with your GitHub or GitLab account.',
                        category: 'Basic Usage'
                    },
                    {
                        title: 'What plugins are used?',
                        value: 'Venus takes 224.7 Earth days to complete one orbit around the Sun.',
                        category: 'Basic Usage'
                    },
                    {
                        title: 'What credentials do I need to use these plugins?',
                        value: 'Binturongs smell like popcorn.',
                        category: 'Basic Usage'
                    },
                    {
                        title: 'How to get Admin privileges?',
                        value: 'Binturongs smell like popcorn.',
                        category: 'Admin Usage'
                    },
                    {
                        title: 'How to set up my own WiseHub?',
                        value: 'Binturongs smell like popcorn.',
                        category: 'Admin Usage'
                    },
                    {
                        title: 'How to write your own WiseHub-Plugin?',
                        value: 'Binturongs smell like popcorn.',
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