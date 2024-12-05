<script lang="ts">
  import Form from "../../../../components/Form.svelte";
  import Main from "../../../../components/Main.svelte";
  import Text from "../../../../components/Text.svelte";
  import Section from "../../../../components/Section.svelte";
  import Select from "../../../../components/Select.svelte";
  import View from "../../../../components/View.svelte";
  import Footer from "../../../../modules_components/Footer.svelte";
  import Header from "../../../../modules_components/Header.svelte";
  import {
    Answer,
    AnswerType,
    AnswerTypes,
    RawAnswer,
    RawAnswerOptions,
    RawAnswerText,
  } from "../../scripts/models/answer";
  import H1 from "../../../../components/H1.svelte";
  import Input from "../../../../components/Input.svelte";
  import Container from "../../../../components/Container.svelte";
  import Button from "../../../../components/Button.svelte";
  import Submit from "../../../../components/Submit.svelte";
  import { AddQuestion } from "./new_question";
  import { RawQuestion } from "../../scripts/models/question";
  import Image from "../../../../components/Image.svelte";
  import Icon from "../../../../components/Icon.svelte";
  import SelectIcon from "../../../../../../public/photos/icons/correct.svg";
  import TrashIcon from "../../../../../../public/photos/icons/trash.svg";
  import Checkbox from "../../../../components/Checkbox.svelte";

  let questionText: string = $state("");
  let answerType: AnswerType = $state(AnswerType.Text);
  let answerDataType: AnswerType = $state(AnswerType.Text);
  let answerData: any = $state(new RawAnswerText());
  $effect(() => {
    if (answerType === answerDataType) return;
    answerDataType = answerType;
    switch (answerDataType - 0) {
      case AnswerType.Options:
        answerData = new RawAnswerOptions();
        break;
      case AnswerType.Text:
        answerData = new RawAnswerText();
        break;
    }
  });

  function refresh(fn: () => any = () => {}): () => any {
    return () => {
      fn();
      answerData = { ...answerData };
    };
  }
</script>

<View>
  <Header />
  <Main changesHeight>
    <Section>
      <Form
        onsubmit={() =>
          AddQuestion(
            new RawQuestion(
              questionText,
              new RawAnswer(answerDataType - 0, answerData),
            ),
          )}
      >
        <H1>Add question</H1>
        <Input bind:value={questionText} label="question" />
        <Select bind:value={answerType} options={AnswerTypes} />
        <Container direction="col">
          {#if answerDataType == AnswerType.Options}
            <Button
              onclick={refresh(
                () => (answerData.answers = [...answerData.answers, ""]),
              )}
              style="add"
            >
              add option
            </Button>

            {#each answerData.answers as _, i}
              <Container direction="row">
                <Input
                  bind:value={answerData.answers[i]}
                  label={i === answerData.correct ? "correct option" : "option"}
                />
                <Button
                  disabled={answerData.correct == i}
                  onclick={refresh(() => (answerData.correct = i))}
                >
                  <Icon src={SelectIcon} alt="select" />
                </Button>
                <Button
                  onclick={refresh(
                    () =>
                      (answerData.answers = answerData.answers.filter(
                        (_, ri) => ri !== i,
                      )),
                  )}
                  style="remove"
                >
                  <Icon src={TrashIcon} alt="remove" />
                </Button>
              </Container>
            {/each}
          {:else if answerDataType == AnswerType.Text}
            <Container direction="row">
              <Checkbox
                bind:value={answerData.case_sensitive}
                onchange={refresh()}
              >
                case sensitive
              </Checkbox>
              <Button
                onclick={refresh(
                  () =>
                    (answerData.correct_answers = [
                      ...answerData.correct_answers,
                      "",
                    ]),
                )}
                style="add"
              >
                add correct answer
              </Button>
            </Container>

            {#each answerData.correct_answers as _, i}
              <Container direction="row">
                <Input
                  bind:value={answerData.correct_answers[i]}
                  label={"correct answer"}
                />
                <Button
                  onclick={refresh(
                    () =>
                      (answerData.correct_answers =
                        answerData.correct_answers.filter((_, ri) => ri !== i)),
                  )}
                  style="remove"
                >
                  <Icon src={TrashIcon} alt="remove" />
                </Button>
              </Container>
            {/each}
          {/if}
          <Submit>Add question</Submit>
        </Container>
      </Form>
    </Section>
  </Main>
  <Footer />
</View>
