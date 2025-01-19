from googletrans import Translator


class AITranslator:
    def __init__(self):
        self.translator = Translator()

    async def translate(self, content: str, lang_from: str, lang_to: str) -> str:
        try:
            translated = await self.translator.translate(content, src=lang_from, dest=lang_to)
            return translated.text
        
        except Exception as e:
            return f"Error: {str(e)}"