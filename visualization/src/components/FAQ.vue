<template>
    <div>
        <b-container fluid="true" class="container-fluid">
               <h2>WiseHub-How-To</h2>
        </b-container>
        <b-container fluid="true" class="container-fluid">
            <div>
            <!-- WITH SCOPED SLOTS and INITIAL OPEN VALUES-->
                <VueFaqAccordion
                        id="accordion"
                        @categorySelect="onCategorySelect"
                        @itemSelect="onItemSelect"
                        :items="myItems"
                        v-slot="itemData"
                        initialTab="Basic Usage"
                        :initialQuestionIndex="1"
                        activeColor="#008B8B"
                >
                    <!-- any html -->
                    <!-- and item data if available via itemData -->
                    <!-- don't add margins to this element (animation become choppy) -->
                    <small>{{ itemData.value }}</small>
                </VueFaqAccordion>
            </div>
        </b-container>
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
                        value: 'Given a 24-hour day and 360 degrees of longitude around the Earth...',
                        category: 'Basic Usage',
                    },
                    {
                        title: 'How to register?',
                        value: 'Venus takes 224.7 Earth days to complete one orbit around the Sun.',
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
                // this.itemData = this.myItems.value
            }
        }
    }
</script>

<style scoped lang="scss">
    .container-fluid {
        width: 75%;
        height: fit-content;
        margin-top: 35px;
        padding: 20px;
        border-radius: 10px;
        background: #F5FFFA;

    }
    #accordion {
        width: 100%;
        padding-top: 30px;
        padding-bottom: 30px;
        margin-left: 80px;
    }
    h2 {
        color: #008B8B;
        width: 100%;
        margin-top: 15px;
        border-radius: 10px;
        text-align: center;
        background: #F5FFFA;
    }
    @media screen and (max-width: 800px) {
        .container-fluid {
            width: 100%;
            padding: 0;
        }
    }


</style>