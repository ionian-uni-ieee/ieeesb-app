module.exports = {
  env: {
    browser: true,
    node: true
  },
  extends: [
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking"
  ],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "tsconfig.json",
    sourceType: "module",
    createDefaultProgram: true,
    ecmaFeatures: {
      jsx: true
    }
  },
  plugins: [
    "@typescript-eslint",
    "@typescript-eslint/tslint",
    "eslint-plugin-import",
    "eslint-plugin-react",
    "eslint-plugin-unicorn",
    "eslint-plugin-prefer-arrow"
  ],
  rules: {
    "@typescript-eslint/interface-name-prefix": "off",
    "react/jsx-indent": [2, 2],
    "react/jsx-indent-props": [2, 2],
    "@typescript-eslint/no-unused-vars": "warn",
    "@typescript-eslint/await-thenable": "error",
    "@typescript-eslint/consistent-type-assertions": "error",
    "@typescript-eslint/member-delimiter-style": [
      "error",
      {
        multiline: {
          delimiter: "none",
          requireLast: true
        },
        singleline: {
          delimiter: "semi",
          requireLast: false
        }
      }
    ],
    semi: ["error", "never"],
    quotes: ["error", "single"],
    "react/jsx-uses-vars": "warn",
    "@typescript-eslint/no-empty-function": "error",
    "@typescript-eslint/no-explicit-any": "off",
    "@typescript-eslint/no-parameter-properties": "off",
    "@typescript-eslint/no-use-before-define": "off",
    "@typescript-eslint/prefer-for-of": "error",
    "@typescript-eslint/prefer-function-type": "error",
    "@typescript-eslint/quotes": ["error", "single"],
    "@typescript-eslint/semi": ["error", "never"],
    "@typescript-eslint/triple-slash-reference": "error",
    "@typescript-eslint/unified-signatures": "error",
    "arrow-body-style": "error",
    "arrow-parens": ["error", "as-needed"],
    camelcase: "error",
    complexity: "off",
    "constructor-super": "error",
    curly: "error",
    "dot-notation": "error",
    "eol-last": "error",
    eqeqeq: ["error", "always"],
    "guard-for-in": "error",
    "id-blacklist": [
      "error",
      "any",
      "Number",
      // "number",
      "String",
      // "string",
      "Boolean",
      // "boolean",
      "Undefined"
      // "undefined"
    ],
    "id-match": "error",
    "import/no-deprecated": "error",
    "linebreak-style": ["error", "unix"],
    "max-classes-per-file": ["error", 1],
    "max-len": [
      "error",
      {
        code: 128
      }
    ],
    "new-parens": "error",
    "no-bitwise": "error",
    "no-caller": "error",
    "no-cond-assign": "error",
    "no-console": "off",
    "no-debugger": "error",
    "no-duplicate-imports": "error",
    "no-empty": "error",
    "no-eval": "error",
    "no-fallthrough": "off",
    "no-invalid-this": "off",
    "no-new-wrappers": "error",
    "no-shadow": [
      "error",
      {
        hoist: "all"
      }
    ],
    "no-throw-literal": "error",
    "no-trailing-spaces": "error",
    "no-undef-init": "error",
    "no-underscore-dangle": "error",
    "no-unsafe-finally": "error",
    "no-unused-expressions": "error",
    "no-unused-labels": "error",
    "no-var": "error",
    "object-shorthand": "error",
    "one-var": ["error", "never"],
    "prefer-arrow/prefer-arrow-functions": "error",
    "prefer-const": "error",
    radix: "error",
    "spaced-comment": "error",
    "use-isnan": "error",
    "valid-typeof": "error",
    "@typescript-eslint/tslint/config": [
      "error",
      {
        rules: {
          "import-spacing": true,
          "jsdoc-format": true,
          "no-reference-import": true,
          "strict-type-predicates": true,
          typedef: true,
          whitespace: true
        }
      }
    ]
  }
};
