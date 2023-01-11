// Generated from java-escape by ANTLR 4.11.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class RiLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.11.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, FUNC=5, INT=6, FLOAT=7, STRING=8, BOOL=9, 
		ANY=10, MUL=11, DIV=12, ADD=13, SUB=14, NUMBER=15, WHITESPACE=16, RETURN=17, 
		IDENTIFIER=18;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "Letter", "LetterOrDigit", "FUNC", "INT", 
			"FLOAT", "STRING", "BOOL", "ANY", "MUL", "DIV", "ADD", "SUB", "NUMBER", 
			"WHITESPACE", "RETURN", "IDENTIFIER"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'{'", "'}'", "'func'", "'int'", "'float'", "'string'", 
			"'bool'", "'any'", "'*'", "'/'", "'+'", "'-'", null, null, "'return'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "FUNC", "INT", "FLOAT", "STRING", "BOOL", 
			"ANY", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "RETURN", 
			"IDENTIFIER"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public RiLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Ri.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0012|\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002"+
		"\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0001\u0000\u0001\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u00046\b\u0004\u0001"+
		"\u0005\u0001\u0005\u0003\u0005:\b\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t"+
		"\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001"+
		"\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u0010\u0004\u0010d\b\u0010\u000b\u0010\f\u0010e\u0001\u0011\u0004\u0011"+
		"i\b\u0011\u000b\u0011\f\u0011j\u0001\u0011\u0001\u0011\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0013\u0001\u0013\u0005\u0013x\b\u0013\n\u0013\f\u0013{\t\u0013\u0000"+
		"\u0000\u0014\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0000\u000b"+
		"\u0000\r\u0005\u000f\u0006\u0011\u0007\u0013\b\u0015\t\u0017\n\u0019\u000b"+
		"\u001b\f\u001d\r\u001f\u000e!\u000f#\u0010%\u0011\'\u0012\u0001\u0000"+
		"\u0006\u0004\u0000$$AZ__az\u0002\u0000\u0000\u007f\u8000\ud800\u8000\udbff"+
		"\u0001\u0000\u8000\ud800\u8000\udbff\u0001\u0000\u8000\udc00\u8000\udfff"+
		"\u0001\u000009\u0003\u0000\t\n\r\r  \u007f\u0000\u0001\u0001\u0000\u0000"+
		"\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001\u0000\u0000"+
		"\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000\u0000"+
		"\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000\u0000"+
		"\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000\u0000"+
		"\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000\u0000"+
		"\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000\u0000"+
		"\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000\u0000"+
		"#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'\u0001"+
		"\u0000\u0000\u0000\u0001)\u0001\u0000\u0000\u0000\u0003+\u0001\u0000\u0000"+
		"\u0000\u0005-\u0001\u0000\u0000\u0000\u0007/\u0001\u0000\u0000\u0000\t"+
		"5\u0001\u0000\u0000\u0000\u000b9\u0001\u0000\u0000\u0000\r;\u0001\u0000"+
		"\u0000\u0000\u000f@\u0001\u0000\u0000\u0000\u0011D\u0001\u0000\u0000\u0000"+
		"\u0013J\u0001\u0000\u0000\u0000\u0015Q\u0001\u0000\u0000\u0000\u0017V"+
		"\u0001\u0000\u0000\u0000\u0019Z\u0001\u0000\u0000\u0000\u001b\\\u0001"+
		"\u0000\u0000\u0000\u001d^\u0001\u0000\u0000\u0000\u001f`\u0001\u0000\u0000"+
		"\u0000!c\u0001\u0000\u0000\u0000#h\u0001\u0000\u0000\u0000%n\u0001\u0000"+
		"\u0000\u0000\'u\u0001\u0000\u0000\u0000)*\u0005(\u0000\u0000*\u0002\u0001"+
		"\u0000\u0000\u0000+,\u0005)\u0000\u0000,\u0004\u0001\u0000\u0000\u0000"+
		"-.\u0005{\u0000\u0000.\u0006\u0001\u0000\u0000\u0000/0\u0005}\u0000\u0000"+
		"0\b\u0001\u0000\u0000\u000016\u0007\u0000\u0000\u000026\b\u0001\u0000"+
		"\u000034\u0007\u0002\u0000\u000046\u0007\u0003\u0000\u000051\u0001\u0000"+
		"\u0000\u000052\u0001\u0000\u0000\u000053\u0001\u0000\u0000\u00006\n\u0001"+
		"\u0000\u0000\u00007:\u0003\t\u0004\u00008:\u0007\u0004\u0000\u000097\u0001"+
		"\u0000\u0000\u000098\u0001\u0000\u0000\u0000:\f\u0001\u0000\u0000\u0000"+
		";<\u0005f\u0000\u0000<=\u0005u\u0000\u0000=>\u0005n\u0000\u0000>?\u0005"+
		"c\u0000\u0000?\u000e\u0001\u0000\u0000\u0000@A\u0005i\u0000\u0000AB\u0005"+
		"n\u0000\u0000BC\u0005t\u0000\u0000C\u0010\u0001\u0000\u0000\u0000DE\u0005"+
		"f\u0000\u0000EF\u0005l\u0000\u0000FG\u0005o\u0000\u0000GH\u0005a\u0000"+
		"\u0000HI\u0005t\u0000\u0000I\u0012\u0001\u0000\u0000\u0000JK\u0005s\u0000"+
		"\u0000KL\u0005t\u0000\u0000LM\u0005r\u0000\u0000MN\u0005i\u0000\u0000"+
		"NO\u0005n\u0000\u0000OP\u0005g\u0000\u0000P\u0014\u0001\u0000\u0000\u0000"+
		"QR\u0005b\u0000\u0000RS\u0005o\u0000\u0000ST\u0005o\u0000\u0000TU\u0005"+
		"l\u0000\u0000U\u0016\u0001\u0000\u0000\u0000VW\u0005a\u0000\u0000WX\u0005"+
		"n\u0000\u0000XY\u0005y\u0000\u0000Y\u0018\u0001\u0000\u0000\u0000Z[\u0005"+
		"*\u0000\u0000[\u001a\u0001\u0000\u0000\u0000\\]\u0005/\u0000\u0000]\u001c"+
		"\u0001\u0000\u0000\u0000^_\u0005+\u0000\u0000_\u001e\u0001\u0000\u0000"+
		"\u0000`a\u0005-\u0000\u0000a \u0001\u0000\u0000\u0000bd\u0007\u0004\u0000"+
		"\u0000cb\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000ec\u0001\u0000"+
		"\u0000\u0000ef\u0001\u0000\u0000\u0000f\"\u0001\u0000\u0000\u0000gi\u0007"+
		"\u0005\u0000\u0000hg\u0001\u0000\u0000\u0000ij\u0001\u0000\u0000\u0000"+
		"jh\u0001\u0000\u0000\u0000jk\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000"+
		"\u0000lm\u0006\u0011\u0000\u0000m$\u0001\u0000\u0000\u0000no\u0005r\u0000"+
		"\u0000op\u0005e\u0000\u0000pq\u0005t\u0000\u0000qr\u0005u\u0000\u0000"+
		"rs\u0005r\u0000\u0000st\u0005n\u0000\u0000t&\u0001\u0000\u0000\u0000u"+
		"y\u0003\t\u0004\u0000vx\u0003\u000b\u0005\u0000wv\u0001\u0000\u0000\u0000"+
		"x{\u0001\u0000\u0000\u0000yw\u0001\u0000\u0000\u0000yz\u0001\u0000\u0000"+
		"\u0000z(\u0001\u0000\u0000\u0000{y\u0001\u0000\u0000\u0000\u0006\u0000"+
		"59ejy\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}