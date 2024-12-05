<script lang="ts">
  import { onMount } from "svelte";
  import H1 from "../../../../components/H1.svelte";
  import Main from "../../../../components/Main.svelte";
  import Section from "../../../../components/Section.svelte";
  import View from "../../../../components/View.svelte";
  import Footer from "../../../../modules_components/Footer.svelte";
  import Header from "../../../../modules_components/Header.svelte";
  import type { RawQuestion } from "../../scripts/models/question";
  import { GetQuestions, RemoveQuestion } from "./get_questions";
  import Text from "../../../../components/Text.svelte";
  import Container from "../../../../components/Container.svelte";
  import Input from "../../../../components/Input.svelte";
  import Checkbox from "../../../../components/Checkbox.svelte";
  import Select from "../../../../components/Select.svelte";
  import Button from "../../../../components/Button.svelte";
  import Article from "../../../../components/Article.svelte";
  import Icon from "../../../../components/Icon.svelte";
  import { getImageSrc } from "../../../user/scripts/value_objects/image_src";

  let questions: RawQuestion[] = $state([]);

  type Searched = "question" | "answer" | "nick";

  let searched: Searched = $state("question");
  let search: string = $state("");
  let case_sensitive: boolean = $state(false);

  let limit: number = $state(50);
  let page: number = $state(0);

  const Search = () => {
    GetQuestions({
      search_question: searched === "question" ? search : "",
      search_question_case_sensitive:
        searched === "question" ? case_sensitive : false,
      search_answer: searched === "answer" ? search : "",
      search_answer_case_sensitive:
        searched === "answer" ? case_sensitive : false,
      search_nick: searched === "nick" ? search : "",
      search_nick_case_sensitive: searched === "nick" ? case_sensitive : false,
      limit: limit,
      page: page,
    }).then((res) => res && (questions = res));
  };

  onMount(Search);
</script>

<View>
  <Header />
  <Main changesHeight>
    <Section>
      <H1>Search questions</H1>

      <Container direction="col">
        <Select
          bind:value={searched}
          options={{
            question: "search in question",
            answer: "search in answer",
            nick: "search in nick",
          }}
        />

        <Container direction="row">
          <Input bind:value={search} label="search question" />
          <Checkbox bind:value={case_sensitive}>case sensitive</Checkbox>
        </Container>

        <Button onclick={Search}>Search</Button>
      </Container>

      {#each questions as question}
        <Article width="fixed">
          <Container direction="row">
            <Icon
              src={getImageSrc(question.creator.user_image, true)}
              alt={question.creator.user_name}
            />
            <Text>{question.question}</Text>
          </Container>
          <!-- <Text>{JSON.stringify(question.answer)}</Text> -->
        </Article>
      {/each}
    </Section>
  </Main>
  <Footer />
</View>
